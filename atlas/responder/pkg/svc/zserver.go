package svc

import (
	models "atlas/responder/internal"
	"atlas/responder/pkg/pb"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"sync"
	"time"
)

const invalidServiceName = "invalid service name"
const emptyRequest = "empty Request"
const success = "success"
const errorMsg = "error"

const (
	// version is the current version of the service
	version = "0.0.1"
)

var (
	// set the environment as instructions.
	pubsubName = os.Getenv("DAPR_PUBSUB_NAME")
	topicName  = "neworder"
)

type Responder struct {
	pb.ResponderServer
	responses *sync.Map
	logger    *logrus.Logger
	responder models.Service
	client    dapr.Client
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

	client, err := dapr.NewClient()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	//defer client.Close()

	return &Responder{
		logger:    logger,
		responses: s,
		responder: getResponder(),
		client:    client,
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

		if err := a.PublishMsg(ctx, id, "getInfo"); err != nil {
			return nil, err
		}

		time.Sleep(time.Millisecond * 15)

		var result interface{}

		a.responses.Range(func(key interface{}, value interface{}) bool {
			if key == id {
				result = value
				return false
			}

			return true
		})

		if result != nil {
			a.responses.Delete(id)
			return &pb.GetInfoResponse{Value: result.(map[string]string)["Value"]}, nil
		}

		return &pb.GetInfoResponse{Value: errorMsg}, nil
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

func (a *Responder) EventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	//log.Printf("event - PubsubName: %responses, Topic: %responses, ID: %responses, Data: %responses", e.PubsubName, e.Topic, e.ID, e.Data)

	log.Println(e.Data)

	var m map[string]string
	json.Unmarshal([]byte(e.Data.(string)), &m)
	log.Println(m["Id"])
	//log.Println(m["Value"])
	a.responses.Store(m["Id"], m)

	return false, nil
}

func (a *Responder) PublishMsg(ctx context.Context, id, command string) error {

	msg := models.Msg{
		Id:      id,
		Command: command,
	}

	data, err := json.Marshal(msg)

	if err != nil {
		return status.Error(codes.Unknown, "err")
	}

	//client, err := dapr.NewClient()
	//if err != nil {
	//	return err
	//}
	//defer client.Close()

	if err := a.client.PublishEvent(ctx, pubsubName, topicName, data); err != nil {
		return err
	}

	return nil
}

func getId(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
