package svc

import (
	models "atlas/portal/internal"
	"atlas/portal/pkg/pb"
	responderpb "atlas/responder/pkg/pb"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

const (
	invalidServiceName = "invalid service name"
	emptyRequest       = "empty Request"
	success            = "success"
	errorMsg           = "error"
	errorMissingResp   = "missing repository response error"
	hiddenUptimeMsg    = "uptime is hidden, mode = false"
	portal             = "portal"
	responder          = "responder"
	storage            = "storage"
)

type Portal struct {
	pb.PortalServer
	client responderpb.ResponderClient
	logger *logrus.Logger
	portal models.Service
}

func getPortal() models.Service {
	return models.Service{
		ServiceDesc:          "portal service desc",
		ServiceUptime:        time.Now().UTC(),
		ServiceCountRequests: 0,
	}
}

func NewPortal(logger *logrus.Logger) (*Portal, error) {
	conn, err := grpc.Dial("0.0.0.0:9091" /*+os.Getenv("PORT")*/, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logger.Fatalf("\nDid not connect %v\n", err)
	}

	client := responderpb.NewResponderClient(conn)

	return &Portal{
		client: client,
		logger: logger,
		portal: getPortal(),
	}, nil
}

func (p *Portal) GetInfo(ctx context.Context, in *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == portal {
		p.portal.Lock()
		p.portal.ServiceCountRequests++
		p.portal.Unlock()
		return &pb.GetInfoResponse{Value: p.portal.ServiceDesc}, nil
	}

	if in.Service == responder || in.Service == storage {
		resp, err := p.client.GetInfo(ctx, &responderpb.GetInfoRequest{Service: in.Service})
		if err != nil {
			return nil, err
		}
		return &pb.GetInfoResponse{Value: resp.Value}, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (p *Portal) SetInfo(ctx context.Context, in *pb.SetInfoRequest) (*pb.SetInfoResponse, error) {
	if in == nil || in.Service == "" || in.Value == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == portal {
		p.portal.Lock()
		p.portal.ServiceCountRequests++
		p.portal.ServiceDesc = in.Value
		p.portal.Unlock()
		return &pb.SetInfoResponse{Msg: success}, nil
	}

	if in.Service == responder || in.Service == storage {
		resp, err := p.client.SetInfo(ctx, &responderpb.SetInfoRequest{Service: in.Service, Value: in.Value})
		if err != nil {
			return nil, err
		}
		return &pb.SetInfoResponse{Msg: resp.Msg}, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (p *Portal) GetUptime(ctx context.Context, in *pb.GetUptimeRequest) (*pb.GetUptimeResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == portal {
		p.portal.Lock()
		p.portal.ServiceCountRequests++
		p.portal.Unlock()
	}

	if in.Service == portal || in.Service == responder || in.Service == storage {
		resp, err := p.client.GetUptime(ctx, &responderpb.GetUptimeRequest{Service: in.Service})
		if err != nil {
			return nil, err
		}

		if in.Service == portal {
			if resp.Value == "true" {
				return &pb.GetUptimeResponse{Value: time.Since(p.portal.ServiceUptime).String()}, nil
			}
			if resp.Value == "false" {
				return &pb.GetUptimeResponse{Value: hiddenUptimeMsg}, nil
			}
		}

		return &pb.GetUptimeResponse{Value: resp.Value}, nil
	}

	return nil, errors.New(invalidServiceName)

}

func (p *Portal) GetRequests(ctx context.Context, in *pb.GetRequestsRequest) (*pb.GetRequestsResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == portal {
		p.portal.Lock()
		p.portal.ServiceCountRequests++
		p.portal.Unlock()
		return &pb.GetRequestsResponse{Value: int32(int(p.portal.ServiceCountRequests))}, nil
	}

	if in.Service == responder || in.Service == storage {
		resp, err := p.client.GetRequests(ctx, &responderpb.GetRequestsRequest{Service: in.Service})
		if err != nil {
			return nil, err
		}
		return &pb.GetRequestsResponse{Value: resp.Value}, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (p *Portal) Reset(ctx context.Context, in *pb.ResetRequest) (*pb.ResetResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == portal {
		p.portal.Lock()
		newPortal := getPortal()
		p.portal.ServiceCountRequests = newPortal.ServiceCountRequests
		p.portal.ServiceDesc = newPortal.ServiceDesc
		p.portal.ServiceUptime = newPortal.ServiceUptime
		p.portal.Unlock()
		return &pb.ResetResponse{Msg: success}, nil
	}

	if in.Service == responder || in.Service == storage {
		resp, err := p.client.Reset(ctx, &responderpb.ResetRequest{Service: in.Service})
		if err != nil {
			return nil, err
		}
		return &pb.ResetResponse{Msg: resp.Msg}, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (p *Portal) GetMode(ctx context.Context, in *pb.GetModeRequest) (*pb.GetModeResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == portal {
		p.portal.Lock()
		p.portal.ServiceCountRequests++
		p.portal.Unlock()
	}

	if in.Service == portal || in.Service == responder || in.Service == storage {
		resp, err := p.client.GetMode(ctx, &responderpb.GetModeRequest{Service: in.Service})
		if err != nil {
			return nil, err
		}
		return &pb.GetModeResponse{Mode: resp.Mode}, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (p *Portal) SetMode(ctx context.Context, in *pb.SetModeRequest) (*pb.SetModeResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == portal {
		p.portal.Lock()
		p.portal.ServiceCountRequests++
		p.portal.Unlock()
	}

	if in.Service == portal || in.Service == responder || in.Service == storage {
		resp, err := p.client.SetMode(ctx, &responderpb.SetModeRequest{Service: in.Service, Mode: in.Mode})
		if err != nil {
			return nil, err
		}
		return &pb.SetModeResponse{Msg: resp.Msg}, nil
	}

	return nil, errors.New(invalidServiceName)
}
