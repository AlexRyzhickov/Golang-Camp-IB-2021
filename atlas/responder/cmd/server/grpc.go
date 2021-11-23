package main

import (
	"atlas/responder/pkg/pb"
	"atlas/responder/pkg/svc"
	"github.com/dapr/go-sdk/service/common"
	dapr "github.com/dapr/go-sdk/service/http"
	"net/http"
	"sync"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/infobloxopen/atlas-app-toolkit/requestid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func NewGRPCServer(logger *logrus.Logger) (*grpc.Server, error) {
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:    time.Duration(viper.GetInt("config.keepalive.time")) * time.Second,
				Timeout: time.Duration(viper.GetInt("config.keepalive.timeout")) * time.Second,
			},
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				// logging middleware
				grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),

				// Request-Id interceptor
				requestid.UnaryServerInterceptor(),

				// Metrics middleware
				grpc_prometheus.UnaryServerInterceptor,

				// validation middleware
				grpc_validator.UnaryServerInterceptor(),

				// collection operators middleware
				gateway.UnaryServerInterceptor(),
			),
		),
	)

	// register service implementation with the grpcServer
	srv, err := svc.NewResponder(logger, &sync.Map{})
	if err != nil {
		return nil, err
	}
	pb.RegisterResponderServer(grpcServer, srv)
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	go func() {
		sub := &common.Subscription{
			PubsubName: "messages2",
			Topic:      "neworder2",
			Route:      "/orders2",
		}

		s := dapr.NewService(":8088")

		if err := s.AddTopicEventHandler(sub, srv.EventHandler); err != nil {
			logger.Fatalf("error adding topic subscription: %v", err)
		}

		if err := s.Start(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("error listenning: %v", err)
		}
	}()

	return grpcServer, nil
}
