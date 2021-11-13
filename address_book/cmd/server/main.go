package main

import (
	"addressbook/internal/config"
	models "addressbook/internal/model"
	"addressbook/internal/pb"
	"addressbook/internal/service"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DBConn), &gorm.Config{})
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
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()

	db, err := connectDB(cfg)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	if err = initializeDB(db); err != nil {
		log.Fatal("failed to init `contact` table", err)
	}

	srv := service.NewAddressBookService(db)

	go func() {
		mux := runtime.NewServeMux()
		if err := pb.RegisterAddressBookServiceHandlerServer(context.Background(), mux, srv); err != nil {
			log.Fatal(err)
		}
		if err = http.ListenAndServe(":"+cfg.ProxyPort, mux); err != nil {
			log.Fatal(err)
		}
	}()

	pb.RegisterAddressBookServiceServer(s, service.NewAddressBookService(db))
	log.Printf("Server is listening on: %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
