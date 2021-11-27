package svc

import (
	models "atlas/responder/internal"
	"atlas/responder/pkg/pb"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	dapr "github.com/dapr/dapr/pkg/proto/runtime/v1"
	"github.com/dapr/go-sdk/service/common"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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
	pubTopic           = "neworder"
	subTopic           = "neworder2"
	route              = "/orders2"
	portal             = "portal"
	responder          = "responder"
	storage            = "storage"
)

type Responder struct {
	pb.ResponderServer
	responses *sync.Map
	logger    *logrus.Logger
	responder models.Service
	client    dapr.DaprClient
	Sub       *common.Subscription
}

func getResponder() models.Service {
	return models.Service{
		ServiceDesc:          "responder service desc",
		ServiceUptime:        time.Now(),
		ServiceCountRequests: 0,
	}
}

func NewResponder(logger *logrus.Logger, s *sync.Map) (*Responder, error) {
	sub := &common.Subscription{
		PubsubName: "messages",
		Topic:      subTopic,
		Route:      route,
	}

	conn, err := grpc.Dial(fmt.Sprintf("0.0.0.0:%s", os.Getenv("DAPR_GRPC_PORT")), grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to open atlas pubsub connection: %v", err)
	}

	return &Responder{
		logger:    logger,
		responses: s,
		responder: getResponder(),
		client:    dapr.NewDaprClient(conn),
		Sub:       sub,
	}, nil
}

func (r *Responder) GetVersion(ctx context.Context, in *empty.Empty) (*pb.VersionResponse, error) {
	return &pb.VersionResponse{Version: version}, nil
}

func (r *Responder) GetInfo(ctx context.Context, in *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == responder {
		r.responder.Lock()
		r.responder.ServiceCountRequests++
		r.responder.Unlock()
		return &pb.GetInfoResponse{Value: r.responder.ServiceDesc}, nil
	}

	if in.Service == storage {
		result, err := r.handleRequest(ctx, "getInfo", nil)
		if err != nil {
			return nil, err
		}
		return &pb.GetInfoResponse{Value: result.(string)}, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (r *Responder) SetInfo(ctx context.Context, in *pb.SetInfoRequest) (*pb.SetInfoResponse, error) {
	if in == nil || in.Service == "" || in.Value == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == responder {
		r.responder.Lock()
		r.responder.ServiceCountRequests++
		r.responder.ServiceDesc = in.Value
		r.responder.Unlock()
		return &pb.SetInfoResponse{Msg: success}, nil
	}

	if in.Service == storage {
		result, err := r.handleRequest(ctx, "setInfo", in.Value)
		if err != nil {
			return nil, err
		}
		return &pb.SetInfoResponse{Msg: result.(string)}, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (r *Responder) GetUptime(ctx context.Context, in *pb.GetUptimeRequest) (*pb.GetUptimeResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == responder {
		r.responder.Lock()
		r.responder.ServiceCountRequests++
		r.responder.Unlock()
	}

	if in.Service == storage || in.Service == responder || in.Service == portal {
		result, err := r.handleRequest(ctx, "getMode", in.Service)
		if err != nil {
			return nil, err
		}
		mode := result.(string)

		if mode == "false" {
			return &pb.GetUptimeResponse{Value: hiddenUptimeMsg}, nil
		}

		if mode == "true" && in.Service == portal {
			return &pb.GetUptimeResponse{Value: mode}, nil
		}

		if mode == "true" && in.Service == responder {
			return &pb.GetUptimeResponse{Value: time.Since(r.responder.ServiceUptime).String()}, nil
		}

		if mode == "true" && in.Service == storage {
			result, err := r.handleRequest(ctx, "getUptime", nil)
			if err != nil {
				return nil, err
			}
			return &pb.GetUptimeResponse{Value: result.(string)}, nil
		}

		return nil, errors.New("invalid value from storage")
	}

	return nil, errors.New(invalidServiceName)
}

func (r *Responder) GetRequests(ctx context.Context, in *pb.GetRequestsRequest) (*pb.GetRequestsResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == responder {
		r.responder.Lock()
		r.responder.ServiceCountRequests++
		r.responder.Unlock()
		return &pb.GetRequestsResponse{Value: int32(int(r.responder.ServiceCountRequests))}, nil
	}

	if in.Service == storage {
		result, err := r.handleRequest(ctx, "getRequests", nil)
		if err != nil {
			return nil, err
		}
		i, err := strconv.Atoi(result.(string))
		if err != nil {
			return nil, err
		}
		return &pb.GetRequestsResponse{Value: int32(i)}, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (r *Responder) Reset(ctx context.Context, in *pb.ResetRequest) (*pb.ResetResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == responder {
		r.responder.Lock()
		newResponder := getResponder()
		r.responder.ServiceCountRequests = newResponder.ServiceCountRequests
		r.responder.ServiceDesc = newResponder.ServiceDesc
		r.responder.ServiceUptime = newResponder.ServiceUptime
		r.responder.Unlock()
		return &pb.ResetResponse{Msg: success}, nil
	}

	if in.Service == storage {
		result, err := r.handleRequest(ctx, "reset", nil)
		if err != nil {
			return nil, err
		}
		return &pb.ResetResponse{Msg: result.(string)}, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (r *Responder) GetMode(ctx context.Context, in *pb.GetModeRequest) (*pb.GetModeResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == responder {
		r.responder.Lock()
		r.responder.ServiceCountRequests++
		r.responder.Unlock()
	}

	if in.Service == storage || in.Service == responder || in.Service == portal {
		result, err := r.handleRequest(ctx, "getMode", in.Service)
		if err != nil {
			return nil, err
		}
		mode := result.(string)
		if mode == "true" || mode == "false" {
			return &pb.GetModeResponse{Mode: mode}, nil
		}
		return nil, errors.New("invalid value from storage")
	}

	return nil, errors.New(invalidServiceName)
}

func (r *Responder) SetMode(ctx context.Context, in *pb.SetModeRequest) (*pb.SetModeResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == responder {
		r.responder.Lock()
		r.responder.ServiceCountRequests++
		r.responder.Unlock()
	}

	if in.Service == storage || in.Service == responder || in.Service == portal {
		result, err := r.handleRequest(ctx, "setMode", in.Service+"&"+strconv.FormatBool(in.Mode))
		if err != nil {
			return nil, err
		}
		return &pb.SetModeResponse{Msg: result.(string)}, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (r *Responder) EventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	r.logger.Debugf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)

	var m map[string]string
	err = json.Unmarshal([]byte(e.Data.(string)), &m)

	if err != nil {
		return false, err
	}

	signal, ok := r.responses.Load(m["Id"])
	if ok {
		signal.(chan interface{}) <- m
	}

	return false, nil
}

func (r *Responder) PublishMsg(ctx context.Context, id, command string, value interface{}) error {

	msg := models.Msg{
		Id:      id,
		Command: command,
		Value:   value,
	}

	data, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	_, err = r.client.PublishEvent(context.Background(), &dapr.PublishEventRequest{
		Topic:      pubTopic,
		Data:       data,
		PubsubName: "messages",
	})

	return err
}

func (r *Responder) handleRequest(ctx context.Context, cmd string, value interface{}) (interface{}, error) {
	id := uuid.NewString()
	if err := r.PublishMsg(ctx, id, cmd, value); err != nil {
		return nil, err
	}

	var result interface{}
	signal := make(chan interface{}, 1)
	r.responses.Store(id, signal)

	for {
		select {
		case result = <-signal:
			r.responses.Delete(id)
			return result.(map[string]string)["Value"], nil
		case <-time.After(5 * time.Second):
			r.responses.Delete(id)
			return nil, errors.New(errorMissingResp)
		}
	}
}
