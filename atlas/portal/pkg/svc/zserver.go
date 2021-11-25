package svc

import (
	models "atlas/portal/internal"
	"atlas/portal/pkg/pb"
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	// version is the current version of the service
	version            = "0.0.1"
	invalidServiceName = "invalid service name"
	emptyRequest       = "empty Request"
	success            = "success"
	errorMsg           = "error"
	errorMissingResp   = "missing repository response error"
	hiddenUptimeMsg    = "uptime is hidden, mode = false"
)

type Portal struct {
	pb.PortalServer
	client pb.PortalClient
	logger *logrus.Logger
	portal models.Service
}

func getPortal() models.Service {
	return models.Service{
		ServiceName:          "portal",
		ServiceDesc:          "portal service desc",
		ServiceUptime:        time.Now(),
		ServiceCountRequests: 0,
	}
}

func NewPortal(logger *logrus.Logger) (*Portal, error) {
	conn, err := grpc.Dial("0.0.0.0:9091" /*+os.Getenv("PORT")*/, grpc.WithInsecure(), grpc.WithBlock())
	log.Println("Hi")
	if err != nil {
		logger.Fatalf("\nDid not connect %v\n", err)
	}

	client := pb.NewPortalClient(conn)

	return &Portal{client: client,
		logger: logger,
		portal: getPortal(),
	}, nil
}

func (p *Portal) GetVersion(context.Context, *empty.Empty) (*pb.VersionResponse, error) {
	return &pb.VersionResponse{Version: version}, nil
}

func (p *Portal) GetInfo(ctx context.Context, in *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == "portal" {
		p.portal.ServiceCountRequests++
		return &pb.GetInfoResponse{Value: p.portal.ServiceDesc}, nil
	}

	if in.Service == "responder" || in.Service == "storage" {
		resp, err := p.client.GetInfo(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (p *Portal) SetInfo(ctx context.Context, in *pb.SetInfoRequest) (*pb.SetInfoResponse, error) {
	if in == nil || in.Service == "" || in.Value == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == "portal" {
		p.portal.ServiceCountRequests++
		p.portal.ServiceDesc = in.Value
		return &pb.SetInfoResponse{Msg: success}, nil
	}

	if in.Service == "responder" || in.Service == "storage" {
		resp, err := p.client.SetInfo(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (p *Portal) GetUptime(ctx context.Context, in *pb.GetUptimeRequest) (*pb.GetUptimeResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == "portal" {
		p.portal.ServiceCountRequests++
	}

	if in.Service == "portal" || in.Service == "responder" || in.Service == "storage" {
		resp, err := p.client.GetUptime(ctx, in)
		if err != nil {
			return nil, err
		}

		if in.Service == "portal" {
			if resp.Value == "true" {
				return &pb.GetUptimeResponse{Value: time.Since(p.portal.ServiceUptime).String()}, nil
			}
			if resp.Value == "false" {
				return &pb.GetUptimeResponse{Value: hiddenUptimeMsg}, nil
			}
		}

		return resp, nil
	}

	return nil, errors.New(invalidServiceName)

}

func (p *Portal) GetRequests(ctx context.Context, in *pb.GetRequestsRequest) (*pb.GetRequestsResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == "portal" {
		p.portal.ServiceCountRequests++
		return &pb.GetRequestsResponse{Value: int32(int(p.portal.ServiceCountRequests))}, nil
	}

	if in.Service == "responder" || in.Service == "storage" {
		resp, err := p.client.GetRequests(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (p *Portal) Reset(ctx context.Context, in *pb.ResetRequest) (*pb.ResetResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == "portal" {
		p.portal.ServiceCountRequests++
		p.portal = getPortal()
		return &pb.ResetResponse{Msg: success}, nil
	}

	if in.Service == "responder" || in.Service == "storage" {
		resp, err := p.client.Reset(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (p *Portal) GetMode(ctx context.Context, in *pb.GetModeRequest) (*pb.GetModeResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == "portal" {
		p.portal.ServiceCountRequests++
	}

	if in.Service == "portal" || in.Service == "responder" || in.Service == "storage" {
		resp, err := p.client.GetMode(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (p *Portal) SetMode(ctx context.Context, in *pb.SetModeRequest) (*pb.SetModeResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == "portal" {
		p.portal.ServiceCountRequests++
	}

	if in.Service == "portal" || in.Service == "responder" || in.Service == "storage" {
		resp, err := p.client.SetMode(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	return nil, errors.New(invalidServiceName)
}
