package service

import (
	"addressbook/internal/pb"
	"context"
	"database/sql"
	"fmt"
	"sync"
)

type AddressBookService struct {
	pb.AddressBookServiceServer
	data sync.Map
	db   *sql.DB
}

func NewAddressBookService(db *sql.DB) *AddressBookService {
	return &AddressBookService{db: db}
}

func (a *AddressBookService) AddContact(_ context.Context, in *pb.AddContactRequest) (*pb.AddContactResponse, error) {
	/*if in == nil || in.Contact == nil {
		return &pb.AddContactResponse{
			Msg: "Contact has not been added",
		}, nil
	}

	a.data.Store(in.Contact.Phone, in.Contact)

	return &pb.AddContactResponse{
		Msg: "Contact added successfully",
	}, nil*/
	return &pb.AddContactResponse{}, nil
}

func (a *AddressBookService) FindContact(_ context.Context, in *pb.FindContactRequest) (*pb.FindContactResponse, error) {
	/*if in == nil || in.Query == "" {
		return &pb.FindContactResponse{
			Msg: "Empty query, contact has not been found",
		}, nil
	}

	contacts := []*pb.Contact{}

	switch in.SearchType {
	case pb.FindContactRequest_NAME:
		a.data.Range(func(key interface{}, value interface{}) bool {
			if contact, ok := value.(*pb.Contact); ok {
				if contact.Name == in.Query {
					contacts = append(contacts, contact)
				}
			}
			return true
		})
	case pb.FindContactRequest_PHONE:
		matchString := in.Query

		a.data.Range(func(key interface{}, value interface{}) bool {
			if contact, ok := value.(*pb.Contact); ok {
				matchString := strings.ReplaceAll(matchString, "?", ".")
				matched, err := regexp.MatchString(matchString, contact.Phone)

				if err == nil {
					if matched {
						contacts = append(contacts, contact)
					}
				}
			}
			return true
		})
	default:
		return &pb.FindContactResponse{
			Msg: "Search value wrong",
		}, nil
	}

	return &pb.FindContactResponse{
		Contacts: contacts,
		Msg:      getFindContactMsg(len(contacts)),
	}, nil*/

	return &pb.FindContactResponse{}, nil
}

func (a *AddressBookService) DeleteContact(_ context.Context, in *pb.DeleteContactRequest) (*pb.DeleteContactResponse, error) {
	/*if in == nil {
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
	}, nil*/
	return &pb.DeleteContactResponse{}, nil
}

func (a *AddressBookService) UpdateContact(_ context.Context, in *pb.UpdateContactRequest) (*pb.UpdateContactResponse, error) {
	/*if in == nil || in.Contact == nil {
		return &pb.UpdateContactResponse{
			Msg: "Contact has not been updated",
		}, nil
	}

	if _, ok := a.data.Load(in.Contact.Phone); !ok {
		return &pb.UpdateContactResponse{
			Msg: "Contact for update not found",
		}, nil
	}

	a.data.Store(in.Contact.Phone, in.Contact)

	return &pb.UpdateContactResponse{
		Msg: "Contact updated successfully",
	}, nil*/
	return &pb.UpdateContactResponse{}, nil
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
