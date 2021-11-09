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

func NewAddressBookService(db *gorm.DB) *AddressBookService {
	return &AddressBookService{db: db}
}

func (a *AddressBookService) AddContact(_ context.Context, in *pb.AddContactRequest) (*pb.AddContactResponse, error) {
	if in == nil || in.Contact == nil {
		return &pb.AddContactResponse{
			Msg: "Contact has not been added",
		}, nil
	}

	err := a.db.Create(&models.Contact{
		Phone:   in.Contact.Phone,
		Name:    in.Contact.Name,
		Address: in.Contact.Address,
	}).Error

	if err != nil {
		return &pb.AddContactResponse{
			Msg: "Сontact has already been added",
		}, err
	}

	return &pb.AddContactResponse{
		Msg: "Contact added successfully",
	}, nil
}

func (a *AddressBookService) FindContact(_ context.Context, in *pb.FindContactRequest) (*pb.FindContactResponse, error) {
	if in == nil || in.Query == "" {
		return &pb.FindContactResponse{
			Msg: "Empty query, contact has not been found",
		}, nil
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
			Msg: "Search value wrong",
		}, nil
	}
}

func (a *AddressBookService) DeleteContact(_ context.Context, in *pb.DeleteContactRequest) (*pb.DeleteContactResponse, error) {
	if in == nil {
		return &pb.DeleteContactResponse{
			Msg: "Contact has not been deleted",
		}, nil
	}

	/*if _, ok := a.data.Load(in.Phone); !ok {
		return &pb.DeleteContactResponse{
			Msg: "Contact not found",
		}, nil
	}

	a.data.Delete(in.Phone)

	return &pb.DeleteContactResponse{
		Msg: "Contact deleted successfully",
	}, nil*/
	return &pb.DeleteContactResponse{}, nil
}

func (a *AddressBookService) UpdateContact(_ context.Context, in *pb.UpdateContactRequest) (*pb.UpdateContactResponse, error) {
	if in == nil || in.Contact == nil {
		return &pb.UpdateContactResponse{
			Msg: "Contact has not been updated",
		}, nil
	}

	/*if _, ok := a.data.Load(in.Contact.Phone); !ok {
		return &pb.UpdateContactResponse{
			Msg: "Contact for update not found",
		}, nil
	}

	a.data.Store(in.Contact.Phone, in.Contact)*/

	return &pb.UpdateContactResponse{
		Msg: "Contact updated successfully",
	}, nil
}

func getFindContactMsg(size int) string {
	if size == 1 {
		return "One contact was found"
	}
	if size > 0 {
		return fmt.Sprintf("Contacts were found successfully, number of contacts: %v", size)
	}
	return "No contacts were found"
}

func processFindContact(findContacts *[]models.Contact, err error) (*pb.FindContactResponse, error) {
	if err != nil {
		return &pb.FindContactResponse{
			Msg: "search error",
		}, err
	}

	contacts := []*pb.Contact{}

	for _, contact := range *findContacts {
		contacts = append(contacts, &pb.Contact{
			Phone:   contact.Phone,
			Name:    contact.Name,
			Address: contact.Address,
		})
	}

	return &pb.FindContactResponse{
		Contacts: contacts,
		Msg:      getFindContactMsg(len(contacts)),
	}, nil
}
