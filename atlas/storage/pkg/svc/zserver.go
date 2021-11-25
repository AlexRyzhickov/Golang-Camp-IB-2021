package svc

import (
	models "atlas/storage/internal/model"
	"context"
	"encoding/json"
	"fmt"
	dapr "github.com/dapr/dapr/pkg/proto/runtime/v1"
	"github.com/dapr/go-sdk/service/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	getInfo         = "getInfo"
	setInfo         = "setInfo"
	getRequests     = "getRequests"
	reset           = "reset"
	getMode         = "getMode"
	setMode         = "setMode"
	getUptime       = "getUptime"
	success         = "success"
	errorMsg        = "invalid commands for storage"
	hiddenUptimeMsg = "uptime is hidden, mode = false"
	pubTopic        = "neworder2"
	subTopic        = "neworder"
	route           = "/orders"
)

type StoragePubSub struct {
	Sub     *common.Subscription
	db      *gorm.DB
	storage models.Service
}

func NewStoragePubSub(db *gorm.DB) *StoragePubSub {
	sub := &common.Subscription{
		PubsubName: "messages",
		Topic:      subTopic,
		Route:      route,
	}

	return &StoragePubSub{
		Sub:     sub,
		db:      db,
		storage: getStorage(),
	}
}

func getStorage() models.Service {
	return models.Service{
		ServiceName:          "storage",
		ServiceDesc:          "storage service desc",
		ServiceUptime:        time.Now(),
		ServiceCountRequests: 0,
	}
}

func (s *StoragePubSub) EventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)
	s.storage.ServiceCountRequests++

	var m map[string]string
	json.Unmarshal([]byte(e.Data.(string)), &m)
	log.Println(m["Id"])

	id := m["Id"]
	command := m["Command"]

	var response interface{}

	switch command {
	case getInfo:
		response = s.storage.ServiceDesc
	case setInfo:
		s.storage.ServiceDesc = m["Value"]
		response = success
	case getRequests:
		response = strconv.Itoa(int(s.storage.ServiceCountRequests))
	case reset:
		s.storage = getStorage()
		response = success
	case getMode:
		serviceName := m["Value"]
		note := models.Note{
			Service: serviceName,
		}
		if err := s.db.First(&note).Error; err != nil {
			return false, err
		}
		response = strconv.FormatBool(note.Mode)
	case setMode:
		values := strings.Split(m["Value"], "&")
		mode, err := strconv.ParseBool(values[1])
		if err != nil {
			return false, err
		}
		note := models.Note{
			Service: values[0],
			Mode:    mode,
		}
		err = s.db.Model(&models.Note{}).Where("service = ?", note.Service).Update("mode", mode).Error
		if err != nil {
			return false, err
		}
		response = success
	case getUptime:
		serviceName := m["Value"]
		note := models.Note{
			Service: serviceName,
		}
		if err := s.db.First(&note).Error; err != nil {
			return false, err
		}
		if note.Mode {
			response = time.Since(s.storage.ServiceUptime).String()
		}
		response = hiddenUptimeMsg
	default:
		response = errorMsg
	}

	if err := PublishMsg(ctx, id, response); err != nil {
		return false, err
	}

	return false, nil
}

func PublishMsg(ctx context.Context, Id string, value interface{}) error {
	response := models.StorageResponse{
		Id:    Id,
		Value: value,
	}

	data, err := json.Marshal(response)

	if err != nil {
		return status.Error(codes.Unknown, "err")
	}

	conn, err := grpc.Dial(fmt.Sprintf("0.0.0.0:%s", os.Getenv("DAPR_GRPC_PORT")), grpc.WithInsecure())
	if err != nil {
		return err
	}

	client := dapr.NewDaprClient(conn)

	_, err = client.PublishEvent(context.Background(), &dapr.PublishEventRequest{
		PubsubName: "messages",
		Topic:      pubTopic,
		Data:       data,
	})

	return nil
}
