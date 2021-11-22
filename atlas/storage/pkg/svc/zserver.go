package svc

import (
	"atlas/storage/pkg/pb"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"gorm.io/gorm"
)

const (
	// version is the current version of the service
	version = "0.0.1"
)

type Storage struct {
	pb.StorageServer
	db *gorm.DB
}

func NewStorage(database *gorm.DB) (*Storage, error) {
	return &Storage{db: database}, nil
}
func (a *Storage) GetVersion(context.Context, *empty.Empty) (*pb.VersionResponse, error) {
	return &pb.VersionResponse{Version: version}, nil
}
