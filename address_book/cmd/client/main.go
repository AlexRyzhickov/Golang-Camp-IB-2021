package main

import (
	"addressbook/internal/pb"
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial(":8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("\nDid not connect %v\n", err)
	}
	defer conn.Close()
	c := pb.NewAddressBookServiceClient(conn)

	log.Println("AddContact block")
	// AddContact
	addResp, err := c.AddContact(ctx, &pb.AddContactRequest{
		Contact: &pb.Contact{
			Name:    "Alex",
			Phone:   "89111726755",
			Address: "Nevsky Prospekt, 11",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(addResp.Msg)

	addResp, err = c.AddContact(ctx, &pb.AddContactRequest{
		Contact: &pb.Contact{
			Name:    "Bob",
			Phone:   "89871726755",
			Address: "Ligovsky Prospekt, 12",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(addResp.Msg)

	addResp, err = c.AddContact(ctx, &pb.AddContactRequest{
		Contact: &pb.Contact{
			Name:    "Mila",
			Phone:   "89872426737",
			Address: "Komendantsky Prospekt, 14",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(addResp.Msg)

	log.Println("FindContact block")
	// FindFindContactByName
	findResp, err := c.FindContactByName(ctx, &pb.FindContactByNameRequest{
		Name: "Alex",
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(findResp.Contacts, findResp.Msg)

	// FindFindContactByPhone
	findResp, err = c.FindContactByPhone(ctx, &pb.FindContactByPhoneRequest{
		Phone: "89871726755",
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(findResp.Contacts, findResp.Msg)

	// FindFindContactByPhone as regular expression
	findResp, err = c.FindContactByPhone(ctx, &pb.FindContactByPhoneRequest{
		Phone: "8987.......",
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(findResp.Contacts, findResp.Msg)

	log.Println("DeleteContact block")
	// DeleteContact
	deleteResp, err := c.DeleteContact(ctx, &pb.DeleteContactRequest{
		Phone: "89871111111",
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(deleteResp.Msg)

	deleteResp, err = c.DeleteContact(ctx, &pb.DeleteContactRequest{
		Phone: "89871726755",
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(deleteResp.Msg)

	log.Println("UpdateContact block")
	// UpdateContact
	updateResp, err := c.UpdateContact(ctx, &pb.UpdateContactRequest{
		Contact: &pb.Contact{
			Name:    "AlexRyzhickov",
			Phone:   "89111726755",
			Address: "Nevsky Prospekt, 11",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(updateResp.Msg)
}
