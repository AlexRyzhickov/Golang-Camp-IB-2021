package svc

import (
	models "atlas/responder/internal"
	"atlas/responder/pkg/pb"
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"time"
)

const invalidServiceName = "invalid service name"
const emptyRequest = "empty Request"
const success = "success"

const (
	// version is the current version of the service
	version = "0.0.1"
)

type Responder struct {
	pb.ResponderServer
	logger    *logrus.Logger
	responder models.Service
}

func getResponder() models.Service {
	return models.Service{
		ServiceName:          "responder",
		ServiceDesc:          "responder service desc",
		ServiceUptime:        time.Now(),
		ServiceCountRequests: 0,
	}
}

func NewResponder(logger *logrus.Logger) (*Responder, error) {

	return &Responder{
		logger:    logger,
		responder: getResponder(),
	}, nil
}

func (a *Responder) GetVersion(ctx context.Context, in *empty.Empty) (*pb.VersionResponse, error) {
	return &pb.VersionResponse{Version: version}, nil
}

func (a *Responder) GetInfo(ctx context.Context, in *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == "responder" {
		a.responder.ServiceCountRequests++
		return &pb.GetInfoResponse{Value: a.responder.ServiceDesc}, nil
	}

	if in.Service == "storage" {

	}

	return nil, errors.New(invalidServiceName)
}

func (a *Responder) SetInfo(ctx context.Context, in *pb.SetInfoRequest) (*pb.SetInfoResponse, error) {
	if in == nil || in.Service == "" || in.Value == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == "responder" {
		a.responder.ServiceDesc = in.Value

		return &pb.SetInfoResponse{Msg: success}, nil
	}

	if in.Service == "storage" {

	}

	return nil, errors.New(invalidServiceName)
}

func (a *Responder) GetUptime(ctx context.Context, in *pb.GetUptimeRequest) (*pb.GetUptimeResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == "responder" {
		return &pb.GetUptimeResponse{
			Value: time.Since(a.responder.ServiceUptime).String(),
		}, nil
	}

	if in.Service == "storage" {

	}

	return nil, errors.New(invalidServiceName)
}

func (a *Responder) GetRequests(ctx context.Context, in *pb.GetRequestsRequest) (*pb.GetRequestsResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == "responder" {
		return &pb.GetRequestsResponse{Value: int32(a.responder.ServiceCountRequests)}, nil
	}

	if in.Service == "storage" {

	}

	return nil, errors.New(invalidServiceName)
}

func (a *Responder) Reset(ctx context.Context, in *pb.ResetRequest) (*pb.ResetResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == "responder" {
		return &pb.ResetResponse{Msg: success}, nil
	}

	if in.Service == "storage" {

	}

	return nil, errors.New(invalidServiceName)
}

func (a *Responder) GetMode(ctx context.Context, in *pb.GetModeRequest) (*pb.GetModeResponse, error) {
	return nil, nil
}

func (a *Responder) SetMode(ctx context.Context, in *pb.GetModeRequest) (*pb.GetModeResponse, error) {
	return nil, nil
}
