package svc

import (
	models "atlas/storage/internal/model"
	"context"
	"encoding/json"
	"fmt"
	dapr "github.com/dapr/dapr/pkg/proto/runtime/v1"
	"github.com/dapr/go-sdk/service/common"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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
	client  dapr.DaprClient
	logger  *logrus.Logger
}

func NewStoragePubSub(db *gorm.DB, logger *logrus.Logger) (*StoragePubSub, error) {
	sub := &common.Subscription{
		PubsubName: "messages",
		Topic:      subTopic,
		Route:      route,
	}

	conn, err := grpc.Dial(fmt.Sprintf("0.0.0.0:%s", os.Getenv("DAPR_GRPC_PORT")), grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to open atlas pubsub connection: %v", err)
	}

	client := dapr.NewDaprClient(conn)

	return &StoragePubSub{
		Sub:     sub,
		db:      db,
		storage: getStorage(),
		client:  client,
		logger:  logger,
	}, nil
}

func getStorage() models.Service {
	return models.Service{
		ServiceDesc:          "storage service desc",
		ServiceUptime:        time.Now().UTC(),
		ServiceCountRequests: 0,
	}
}

func (s *StoragePubSub) EventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	s.logger.Debugf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)

	s.storage.Lock()
	s.storage.ServiceCountRequests++
	s.storage.Unlock()

	var m map[string]string
	json.Unmarshal([]byte(e.Data.(string)), &m)

	id := m["Id"]
	command := m["Command"]

	var response interface{}

	switch command {
	case getInfo:
		response = s.storage.ServiceDesc
	case setInfo:
		s.storage.Lock()
		s.storage.ServiceDesc = m["Value"]
		s.storage.Unlock()
		response = success
	case getRequests:
		response = strconv.Itoa(int(s.storage.ServiceCountRequests))
	case reset:
		s.storage.Lock()
		newStorage := getStorage()
		s.storage.ServiceCountRequests = newStorage.ServiceCountRequests
		s.storage.ServiceUptime = newStorage.ServiceUptime
		s.storage.ServiceDesc = newStorage.ServiceDesc
		s.storage.Unlock()
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
		note := models.Note{
			Service: "storage",
		}
		if err := s.db.First(&note).Error; err != nil {
			return false, err
		}
		if note.Mode {
			response = time.Since(s.storage.ServiceUptime).String()
		} else {
			response = hiddenUptimeMsg
		}
	default:
		response = errorMsg
	}

	if err := s.PublishMsg(ctx, id, response); err != nil {
		return false, err
	}

	return false, nil
}

func (s *StoragePubSub) PublishMsg(ctx context.Context, Id string, value interface{}) error {
	response := models.StorageResponse{
		Id:    Id,
		Value: value,
	}

	data, err := json.Marshal(response)

	if err != nil {
		return status.Error(codes.Unknown, "err")
	}

	_, err = s.client.PublishEvent(context.Background(), &dapr.PublishEventRequest{
		PubsubName: "messages",
		Topic:      pubTopic,
		Data:       data,
	})

	return nil
}
