package svc

import (
	models "atlas/responder/internal"
	"atlas/responder/pkg/pb"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	dapr "github.com/dapr/dapr/pkg/proto/runtime/v1"
	"github.com/dapr/go-sdk/service/common"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		ServiceName:          "responder",
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
		id := getId(time.Now().String())
		if err := a.PublishMsg(ctx, id, "getInfo", nil); err != nil {
			return nil, err
		}
		time.Sleep(time.Millisecond * 15)
		result, ok := a.responses.Load(id)
		if !ok {
			return &pb.GetInfoResponse{Value: errorMissingResp}, nil
		}
		a.responses.Delete(id)
		return &pb.GetInfoResponse{Value: result.(map[string]string)["Value"]}, nil
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
		id := getId(time.Now().String())
		if err := a.PublishMsg(ctx, id, "setInfo", in.Value); err != nil {
			return nil, err
		}
		time.Sleep(time.Millisecond * 15)
		result, ok := a.responses.Load(id)
		if !ok {
			return &pb.SetInfoResponse{Msg: errorMissingResp}, nil
		}
		a.responses.Delete(id)
		return &pb.SetInfoResponse{Msg: result.(map[string]string)["Value"]}, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (a *Responder) GetUptime(ctx context.Context, in *pb.GetUptimeRequest) (*pb.GetUptimeResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == "storage" || in.Service == "responder" || in.Service == "portal" {
		id := getId(time.Now().String())
		if err := a.PublishMsg(ctx, id, "getMode", in.Service); err != nil {
			return nil, err
		}
		time.Sleep(time.Millisecond * 15)
		result, ok := a.responses.Load(id)
		if !ok {
			return nil, errors.New(errorMissingResp)
		}
		a.responses.Delete(id)
		mode := result.(map[string]string)["Value"]

		if mode == "false" {
			return &pb.GetUptimeResponse{Value: hiddenUptimeMsg}, nil
		}

		if mode == "true" && in.Service == "responder" {
			return &pb.GetUptimeResponse{Value: time.Since(a.responder.ServiceUptime).String()}, nil
		}

		if mode == "true" && in.Service == "storage" {
			id := getId(time.Now().String())
			if err := a.PublishMsg(ctx, id, "getUptime", nil); err != nil {
				return nil, err
			}
			time.Sleep(time.Millisecond * 15)
			result, ok := a.responses.Load(id)
			if !ok {
				return nil, errors.New(errorMissingResp)
			}
			a.responses.Delete(id)
			return &pb.GetUptimeResponse{Value: result.(map[string]string)["Value"]}, nil
		}

		return nil, errors.New("invalid value from storage")
	}

	return nil, errors.New(invalidServiceName)
}

func (a *Responder) GetRequests(ctx context.Context, in *pb.GetRequestsRequest) (*pb.GetRequestsResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == "responder" {
		a.responder.ServiceCountRequests++
		return &pb.GetRequestsResponse{Value: int32(int(a.responder.ServiceCountRequests))}, nil
	}

	if in.Service == "storage" {
		id := getId(time.Now().String())
		if err := a.PublishMsg(ctx, id, "getRequests", nil); err != nil {
			return nil, err
		}
		time.Sleep(time.Millisecond * 15)
		result, ok := a.responses.Load(id)
		if !ok {
			return nil, errors.New(errorMissingResp)
		}
		a.responses.Delete(id)
		i, err := strconv.Atoi(result.(map[string]string)["Value"])
		if err != nil {
			return nil, err
		}
		return &pb.GetRequestsResponse{Value: int32(i)}, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (a *Responder) Reset(ctx context.Context, in *pb.ResetRequest) (*pb.ResetResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == "responder" {
		a.responder = getResponder()
		return &pb.ResetResponse{Msg: success}, nil
	}

	if in.Service == "storage" {
		id := getId(time.Now().String())
		if err := a.PublishMsg(ctx, id, "reset", nil); err != nil {
			return nil, err
		}
		time.Sleep(time.Millisecond * 15)
		result, ok := a.responses.Load(id)
		if !ok {
			return nil, errors.New(errorMissingResp)
		}
		a.responses.Delete(id)
		return &pb.ResetResponse{Msg: result.(map[string]string)["Value"]}, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (a *Responder) GetMode(ctx context.Context, in *pb.GetModeRequest) (*pb.GetModeResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == "storage" || in.Service == "responder" || in.Service == "portal" {
		id := getId(time.Now().String())
		if err := a.PublishMsg(ctx, id, "getMode", in.Service); err != nil {
			return nil, err
		}
		time.Sleep(time.Millisecond * 15)
		result, ok := a.responses.Load(id)
		if !ok {
			return nil, errors.New(errorMissingResp)
		}
		a.responses.Delete(id)
		mode := result.(map[string]string)["Value"]

		if mode == "true" || mode == "false" {
			return &pb.GetModeResponse{Mode: mode}, nil
		}

		return nil, errors.New("invalid value from storage")
	}

	return nil, errors.New(invalidServiceName)
}

func (a *Responder) SetMode(ctx context.Context, in *pb.SetModeRequest) (*pb.SetModeResponse, error) {
	if in == nil || in.Service == "" {
		return nil, errors.New(emptyRequest)
	}

	if in.Service == "storage" || in.Service == "responder" || in.Service == "portal" {
		id := getId(time.Now().String())
		if err := a.PublishMsg(ctx, id, "setMode", in.Service+"&"+strconv.FormatBool(in.Mode)); err != nil {
			return nil, err
		}
		time.Sleep(time.Millisecond * 15)
		result, ok := a.responses.Load(id)
		if !ok {
			return nil, errors.New(errorMissingResp)
		}
		a.responses.Delete(id)
		return &pb.SetModeResponse{Msg: result.(map[string]string)["Value"]}, nil
	}

	return nil, errors.New(invalidServiceName)
}

func (a *Responder) EventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)

	var m map[string]string
	err = json.Unmarshal([]byte(e.Data.(string)), &m)

	if err != nil {
		return false, err
	}

	a.responses.Store(m["Id"], m)

	return false, nil
}

func (a *Responder) PublishMsg(ctx context.Context, id, command string, value interface{}) error {

	msg := models.Msg{
		Id:      id,
		Command: command,
		Value:   value,
	}

	data, err := json.Marshal(msg)

	if err != nil {
		return status.Error(codes.Unknown, "err")
	}

	_, err = a.client.PublishEvent(context.Background(), &dapr.PublishEventRequest{
		Topic:      pubTopic,
		Data:       data,
		PubsubName: "messages",
	})

	return err
}

func getId(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
