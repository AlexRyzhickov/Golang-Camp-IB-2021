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

func (a *Responder) GetInfo(context.Context, *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	return nil, nil
}

func (a *Responder) SetInfo(context.Context, *pb.SetInfoRequest) (*pb.SetInfoResponse, error) {
	return nil, nil
}

func (a *Responder) GetUptime(context.Context, *pb.GetUptimeRequest) (*pb.GetUptimeResponse, error) {
	return nil, nil
}

func (a *Responder) GetRequests(context.Context, *pb.GetRequestsRequest) (*pb.GetRequestsResponse, error) {
	return nil, nil
}

func (a *Responder) Reset(context.Context, *pb.ResetRequest) (*pb.ResetResponse, error) {
	return nil, nil
}

func (a *Responder) GetMode(context.Context, *pb.GetModeRequest) (*pb.GetModeResponse, error) {
	return nil, nil
}

func (a *Responder) SetMode(context.Context, *pb.GetModeRequest) (*pb.GetModeResponse, error) {
	return nil, nil
}
