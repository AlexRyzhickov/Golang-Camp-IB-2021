package svc

import (
	"atlas/portal/pkg/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

const (
	// version is the current version of the service
	version = "0.0.1"
)

type Portal struct {
	pb.PortalServer
}

func NewPortal() (*Portal, error) {
	return &Portal{}, nil
}

func (a *Portal) GetVersion(context.Context, *empty.Empty) (*pb.VersionResponse, error) {
	return &pb.VersionResponse{Version: version}, nil
}

func (a *Portal) GetInfo(context.Context, *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	return nil, nil
}

func (a *Portal) SetInfo(context.Context, *pb.SetInfoRequest) (*pb.SetInfoResponse, error) {
	return nil, nil
}

func (a *Portal) GetUptime(context.Context, *pb.GetUptimeRequest) (*pb.GetUptimeResponse, error) {
	return nil, nil
}

func (a *Portal) GetRequests(context.Context, *pb.GetRequestsRequest) (*pb.GetRequestsResponse, error) {
	return nil, nil
}

func (a *Portal) Reset(context.Context, *pb.ResetRequest) (*pb.ResetResponse, error) {
	return nil, nil
}

func (a *Portal) GetMode(context.Context, *pb.GetModeRequest) (*pb.GetModeResponse, error) {
	return nil, nil
}

func (a *Portal) SetMode(context.Context, *pb.GetModeRequest) (*pb.GetModeResponse, error) {
	return nil, nil
}
