package main

import (
	models "addressbook/internal/model"
	"addressbook/internal/pb"
	"addressbook/internal/service"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "postgres://postgres:postgres@db/backend?sslmode=disable"

//const dsn = "postgres://postgres:postgres@db user=postgres password=postgres dbname=backend port=5432 sslmode=disable"

func connectDB() (*gorm.DB, error) {
	dsn := dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initializeDB(db *gorm.DB) error {
	if isInit := db.Migrator().HasTable(&models.Contact{}); !isInit {
		err := db.Migrator().CreateTable(&models.Contact{})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {

	lis, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()

	db, err := connectDB()
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	if err = initializeDB(db); err != nil {
		log.Fatal("failed to init `contact` table", err)
	}

	pb.RegisterAddressBookServiceServer(s, service.NewAddressBookService(db))
	log.Printf("Server is listening on: %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
