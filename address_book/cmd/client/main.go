package main

import (
	"addressbook/internal/pb"
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial(":"+os.Getenv("PORT"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("\nDid not connect %v\n", err)
	}
	defer conn.Close()
	c := pb.NewAddressBookServiceClient(conn)

	// AddContact
	{
		log.Println("AddContact block")
		addResp, err := c.AddContact(ctx, &pb.AddContactRequest{
			Contact: &pb.Contact{
				Name:    "Alex",
				Phone:   "89111726755",
				Address: "Nevsky Prospekt, 11",
			},
		})
		if err != nil {
			log.Println(err)
		} else {
			log.Println(addResp.Msg)
		}

		addResp, err = c.AddContact(ctx, &pb.AddContactRequest{
			Contact: &pb.Contact{
				Name:    "Bob",
				Phone:   "89871726755",
				Address: "Ligovsky Prospekt, 12",
			},
		})
		if err != nil {
			log.Println(err)
		} else {
			log.Println(addResp.Msg)
		}

		addResp, err = c.AddContact(ctx, &pb.AddContactRequest{
			Contact: &pb.Contact{
				Name:    "Mila",
				Phone:   "89872426737",
				Address: "Komendantsky Prospekt, 14",
			},
		})
		if err != nil {
			log.Println(err)
		} else {
			log.Println(addResp.Msg)
		}
	}

	// FindContact
	{
		log.Println("FindContact block")
		// FindFindContactByName
		findResp, err := c.FindContact(ctx, &pb.FindContactRequest{
			Query:      "Alex",
			SearchType: pb.FindContactRequest_NAME,
		})

		if err != nil {
			log.Fatal(err)
		}
		log.Println(findResp.Contacts, findResp.Msg)

		// FindFindContactByPhone
		findResp, err = c.FindContact(ctx, &pb.FindContactRequest{
			Query:      "89871726755",
			SearchType: pb.FindContactRequest_PHONE,
		})

		if err != nil {
			log.Fatal(err)
		}
		log.Println(findResp.Contacts, findResp.Msg)

		// FindFindContactByPhone as wildcards
		findResp, err = c.FindContact(ctx, &pb.FindContactRequest{
			Query:      "8987???????",
			SearchType: pb.FindContactRequest_PHONE,
		})

		if err != nil {
			log.Fatal(err)
		}
		log.Println(findResp.Contacts, findResp.Msg)

		// FindFindContactByPhone as wildcards
		findResp, err = c.FindContact(ctx, &pb.FindContactRequest{
			Query:      "8987*",
			SearchType: pb.FindContactRequest_PHONE,
		})

		if err != nil {
			log.Fatal(err)
		}
		log.Println(findResp.Contacts, findResp.Msg)
	}

	// DeleteContact
	{
		log.Println("DeleteContact block")

		deleteResp, err := c.DeleteContact(ctx, &pb.DeleteContactRequest{
			Phone: "89871111111",
		})

		if err != nil {
			log.Println(err)
		} else {
			log.Println(deleteResp.Msg)
		}

		deleteResp, err = c.DeleteContact(ctx, &pb.DeleteContactRequest{
			Phone: "89871726755",
		})

		if err != nil {
			log.Println(err)
		} else {
			log.Println(deleteResp.Msg)
		}
	}

	// UpdateContact
	{
		log.Println("UpdateContact block")
		updateResp, err := c.UpdateContact(ctx, &pb.UpdateContactRequest{
			Contact: &pb.Contact{
				Phone:   "89111726755",
				Name:    "AlexRyzhickov",
				Address: "Nevsky Prospekt, 11",
			},
		})
		if err != nil {
			log.Println(err)
		} else {
			log.Println(updateResp.Msg)
		}
	}
}
