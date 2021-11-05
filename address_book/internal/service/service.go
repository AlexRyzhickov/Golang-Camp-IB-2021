package service

import (
	"addressbook/internal/pb"
	"context"
	"fmt"
	"regexp"
	"sync"
)

type AddressBookService struct {
	pb.AddressBookServiceServer
	data sync.Map
}

func (a *AddressBookService) AddContact(_ context.Context, in *pb.AddContactRequest) (*pb.AddContactResponse, error) {
	if in == nil || in.Contact == nil {
		return &pb.AddContactResponse{
			Msg: "Contact has not been added",
		}, nil
	}

	a.data.Store(in.Contact.Phone, in.Contact)

	return &pb.AddContactResponse{
		Msg: "Contact added successfully",
	}, nil
}

func (a *AddressBookService) FindContactByName(_ context.Context, in *pb.FindContactByNameRequest) (*pb.FindContactResponse, error) {
	if in == nil {
		return &pb.FindContactResponse{
			Msg: "Contact has not been found",
		}, nil
	}

	contacts := []*pb.Contact{}

	a.data.Range(func(key interface{}, value interface{}) bool {
		if contact, ok := value.(*pb.Contact); ok {
			if contact.Name == in.Name {
				contacts = append(contacts, contact)
			}
		}
		return true
	})

	return &pb.FindContactResponse{
		Contacts: contacts,
		Msg:      getFindContactMsg(len(contacts)),
	}, nil
}

func (a *AddressBookService) FindContactByPhone(_ context.Context, in *pb.FindContactByPhoneRequest) (*pb.FindContactResponse, error) {
	if in == nil {
		return &pb.FindContactResponse{
			Msg: "Contact has not been found",
		}, nil
	}

	matchString := in.Phone
	contacts := []*pb.Contact{}

	a.data.Range(func(key interface{}, value interface{}) bool {
		if contact, ok := value.(*pb.Contact); ok {
			matched, err := regexp.MatchString(matchString, contact.Phone)

			if err == nil {
				if matched {
					contacts = append(contacts, contact)
				}
			}
		}
		return true
	})

	return &pb.FindContactResponse{
		Contacts: contacts,
		Msg:      getFindContactMsg(len(contacts)),
	}, nil
}

func (a *AddressBookService) DeleteContact(_ context.Context, in *pb.DeleteContactRequest) (*pb.DeleteContactResponse, error) {
	if in == nil {
		return &pb.DeleteContactResponse{
			Msg: "Contact has not been deleted",
		}, nil
	}

	if _, ok := a.data.Load(in.Phone); !ok {
		return &pb.DeleteContactResponse{
			Msg: "Contact not found",
		}, nil
	}

	a.data.Delete(in.Phone)

	return &pb.DeleteContactResponse{
		Msg: "Contact deleted successfully",
	}, nil
}

func (a *AddressBookService) UpdateContact(_ context.Context, in *pb.UpdateContactRequest) (*pb.UpdateContactResponse, error) {
	if in == nil || in.Contact == nil {
		return &pb.UpdateContactResponse{
			Msg: "Contact has not been updated",
		}, nil
	}

	a.data.Store(in.Contact.Phone, in.Contact)

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
