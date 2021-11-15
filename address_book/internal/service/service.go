package service

import (
	models "addressbook/internal/model"
	"addressbook/internal/pb"
	"context"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

const invalidInputData = "Invalid input data error"
const successAdding = "Contact added successfully"
const successDeleting = "Contact deleted successfully"
const successUpdating = "Contact updated successfully"
const addError = "Adding contact error"
const updateError = "Updating contact error"
const deleteError = "Deleting contact error"
const findError = "Search contact error"
const wrongValueSearchError = "Search value wrong error"
const findZeroContactMsg = "No contacts were found"
const findOneContactMsg = "One contact was found successfully"
const findSomeContactsMsg = "Contacts were found successfully, number of contacts:"

type DBInterface interface {
	FirstOrCreate(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
	Model(value interface{}) (tx *gorm.DB)
	Updates(values interface{}) (tx *gorm.DB)
}

type AddressBookService struct {
	pb.AddressBookServiceServer
	db DBInterface
}

func NewAddressBookService(db *gorm.DB) *AddressBookService {
	return &AddressBookService{db: db}
}

func (a *AddressBookService) AddContact(_ context.Context, in *pb.AddContactRequest) (*pb.AddContactResponse, error) {
	if in == nil || in.Contact == nil {
		return &pb.AddContactResponse{Msg: invalidInputData}, nil
	}

	err := a.db.FirstOrCreate(&models.Contact{
		Phone:   in.Contact.Phone,
		Name:    in.Contact.Name,
		Address: in.Contact.Address,
	}).Error

	if err != nil {
		return &pb.AddContactResponse{Msg: fmt.Sprintf("%s, %v", addError, err.Error())}, err
	}

	return &pb.AddContactResponse{Msg: successAdding}, nil
}

func (a *AddressBookService) FindContact(_ context.Context, in *pb.FindContactRequest) (*pb.FindContactResponse, error) {
	if in == nil || in.Query == "" {
		return &pb.FindContactResponse{Msg: invalidInputData}, nil
	}

	switch in.SearchType {
	case pb.FindContactRequest_NAME:
		findContacts := []models.Contact{}

		err := a.db.Where("name = ?", in.Query).Find(&findContacts).Error

		return processFindContact(&findContacts, err)
	case pb.FindContactRequest_PHONE:
		matchString := strings.ReplaceAll(in.Query, "?", "_")
		matchString = strings.ReplaceAll(matchString, "*", "%")

		findContacts := []models.Contact{}

		err := a.db.Where("phone LIKE ?", matchString).Find(&findContacts).Error

		return processFindContact(&findContacts, err)
	default:
		return &pb.FindContactResponse{
			Msg: wrongValueSearchError,
		}, nil
	}
}

func (a *AddressBookService) DeleteContact(_ context.Context, in *pb.DeleteContactRequest) (*pb.DeleteContactResponse, error) {
	if in == nil {
		return &pb.DeleteContactResponse{Msg: invalidInputData}, nil
	}

	err := a.db.Delete(&models.Contact{}, in.Phone).Error

	if err != nil {
		return &pb.DeleteContactResponse{Msg: fmt.Sprintf("%s, %v", deleteError, err)}, err
	}

	return &pb.DeleteContactResponse{Msg: successDeleting}, nil
}

func (a *AddressBookService) UpdateContact(_ context.Context, in *pb.UpdateContactRequest) (*pb.UpdateContactResponse, error) {
	if in == nil || in.Contact == nil {
		return &pb.UpdateContactResponse{
			Msg: invalidInputData,
		}, nil
	}

	contact := models.Contact{
		Phone:   in.Contact.Phone,
		Name:    in.Contact.Name,
		Address: in.Contact.Address,
	}

	err := a.db.Model(&contact).Updates(models.Contact{Name: contact.Name, Address: contact.Address}).Error

	if err != nil {
		return &pb.UpdateContactResponse{Msg: fmt.Sprintf("%s, %v", updateError, err)}, err
	}

	return &pb.UpdateContactResponse{Msg: successUpdating}, nil
}

func getFindContactMsg(size int) string {
	if size == 1 {
		return findOneContactMsg
	}
	if size > 0 {
		return fmt.Sprintf("%s %v", findSomeContactsMsg, size)
	}
	return findZeroContactMsg
}

func processFindContact(findContacts *[]models.Contact, err error) (*pb.FindContactResponse, error) {
	if err != nil {
		return &pb.FindContactResponse{Msg: fmt.Sprintf("%s, %v", findError, err)}, err
	}

	contacts := []*pb.Contact{}

	for _, contact := range *findContacts {
		contacts = append(contacts, &pb.Contact{
			Phone:   contact.Phone,
			Name:    contact.Name,
			Address: contact.Address,
		})
	}

	return &pb.FindContactResponse{Contacts: contacts, Msg: getFindContactMsg(len(contacts))}, nil
}
