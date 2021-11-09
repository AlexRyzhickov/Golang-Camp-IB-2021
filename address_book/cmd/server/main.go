package main

import (
	"addressbook/internal/pb"
	"addressbook/internal/service"
	"database/sql"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const DBConnString = "postgres://postgres:postgres@127.0.0.1:5432/backend?sslmode=disable"

func main() {

	lis, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()

	db, err := sql.Open("postgres", DBConnString)
	if err != nil {
		log.Fatal("DB initializing error", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("DB pinging error", err)
	}

	pb.RegisterAddressBookServiceServer(s, service.NewAddressBookService(db))
	log.Printf("Server is listening on: %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
