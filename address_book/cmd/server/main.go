package main

import (
	"addressbook/internal/pb"
	"addressbook/internal/service"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()

	pb.RegisterAddressBookServiceServer(s, &service.AddressBookService{})
	log.Printf("Server is listening on: %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
