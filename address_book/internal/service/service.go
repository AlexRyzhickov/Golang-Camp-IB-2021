package service

import (
	models "addressbook/internal/model"
	"addressbook/internal/pb"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type AddressBookService struct {
	pb.AddressBookServiceServer
	db *gorm.DB
}

const invalidInputData = "Invalid input data error"
const successAdding = "Contact added successfully"
const successDeleting = "Contact deleted successfully"
const successUpdating = "Contact updated successfully"
const addError = "Adding contact error"
const updateError = "Updating contact error"
const deleteError = "Deleting contact error"
const findError = "Search contact error"
const wrongValueSearchError = "Search value wrong error"
const addingDuplicateContactErr = "Contact has already been added"

func NewAddressBookService(db *gorm.DB) *AddressBookService {
	return &AddressBookService{db: db}
}

func (a *AddressBookService) AddContact(_ context.Context, in *pb.AddContactRequest) (*pb.AddContactResponse, error) {
	if in == nil || in.Contact == nil {
		return &pb.AddContactResponse{Msg: invalidInputData}, nil
	}

	var count int64 = 0
	if err := a.db.Limit(1).Find(&models.Contact{}, in.Contact.Phone).Count(&count).Error; err != nil {
		return &pb.AddContactResponse{Msg: addError}, nil
	}

	if count == 1 {
		return &pb.AddContactResponse{Msg: addingDuplicateContactErr}, nil
	}

	err := a.db.Create(&models.Contact{
		Phone:   in.Contact.Phone,
		Name:    in.Contact.Name,
		Address: in.Contact.Address,
	}).Error

	if err != nil {
		return &pb.AddContactResponse{Msg: fmt.Sprintf("%s %v", addError, err)}, err
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
		findContacts := []models.Contact{}

		err := a.db.Where("phone LIKE ?", in.Query).Find(&findContacts).Error

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
		return &pb.DeleteContactResponse{Msg: fmt.Sprintf("%s %v", deleteError, err)}, nil
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
		return &pb.UpdateContactResponse{Msg: fmt.Sprintf("%s %v", updateError, err)}, nil
	}

	return &pb.UpdateContactResponse{Msg: successUpdating}, nil
}

func getFindContactMsg(size int) string {
	if size == 1 {
		return "One contact was found successfully"
	}
	if size > 0 {
		return fmt.Sprintf("Contacts were found successfully, number of contacts: %v", size)
	}
	return "No contacts were found"
}

func processFindContact(findContacts *[]models.Contact, err error) (*pb.FindContactResponse, error) {
	if err != nil {
		return &pb.FindContactResponse{Msg: fmt.Sprintf("%s %v", findError, err)}, err
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
