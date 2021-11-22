package svc

import (
	"atlas/responder/pkg/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

const (
	// version is the current version of the service
	version = "0.0.1"
)

type Responder struct {
	pb.ResponderServer
}

func NewResponder() (*Responder, error) {
	return &Responder{}, nil
}
func (a *Responder) GetVersion(context.Context, *empty.Empty) (*pb.VersionResponse, error) {
	return &pb.VersionResponse{Version: version}, nil
}
