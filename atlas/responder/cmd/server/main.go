package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/infobloxopen/atlas-app-toolkit/gorm/resource"
	"github.com/infobloxopen/atlas-app-toolkit/server"
	pubsubgrpc "github.com/infobloxopen/atlas-pubsub/grpc"
)

func main() {
	doneC := make(chan error)
	logger := NewLogger()
	if viper.GetBool("internal.enable") {
		go func() { doneC <- ServeInternal(logger) }()
	}

	go func() { doneC <- ServeExternal(logger) }()

	if viper.GetBool("atlas.pubsub.enable") {
		InitSubscriber(logger)
	}

	if err := <-doneC; err != nil {
		logger.Fatal(err)
	}
}

func NewLogger() *logrus.Logger {
	logger := logrus.StandardLogger()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(true)

	// Set the log level on the default logger based on command line flag
	if level, err := logrus.ParseLevel(viper.GetString("logging.level")); err != nil {
		logger.Errorf("Invalid %q provided for log level", viper.GetString("logging.level"))
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(level)
	}

	return logger
}

// ServeInternal builds and runs the server that listens on InternalAddress
func ServeInternal(logger *logrus.Logger) error {

	s, err := server.NewServer(
		// register metrics
		server.WithHandler("/metrics", promhttp.Handler()),
	)
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", viper.GetString("internal.address"), viper.GetString("internal.port")))
	if err != nil {
		return err
	}

	logger.Debugf("serving internal http at %q", fmt.Sprintf("%s:%s", viper.GetString("internal.address"), viper.GetString("internal.port")))
	return s.Serve(nil, l)
}

// ServeExternal builds and runs the server that listens on ServerAddress and GatewayAddress
func ServeExternal(logger *logrus.Logger) error {

	grpcServer, err := NewGRPCServer(logger)
	if err != nil {
		logger.Fatalln(err)
	}
	grpc_prometheus.Register(grpcServer)

	s, err := server.NewServer(
		server.WithGrpcServer(grpcServer),
	)
	if err != nil {
		logger.Fatalln(err)
	}

	grpcL, err := net.Listen("tcp", fmt.Sprintf("%s:%s", viper.GetString("server.address"), viper.GetString("server.port")))
	if err != nil {
		logger.Fatalln(err)
	}

	logger.Printf("serving gRPC at %s:%s", viper.GetString("server.address"), viper.GetString("server.port"))

	return s.Serve(grpcL, nil)
}

func init() {
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AddConfigPath(viper.GetString("config.source"))
	if viper.GetString("config.file") != "" {
		log.Printf("Serving from configuration file: %s", viper.GetString("config.file"))
		viper.SetConfigName(viper.GetString("config.file"))
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("cannot load configuration: %v", err)
		}
	} else {
		log.Printf("Serving from default values, environment variables, and/or flags")
	}
	resource.RegisterApplication(viper.GetString("app.id"))
	resource.SetPlural()
}

// InitSubscriber initiliazes the example atlas-pubsub subscriber
func InitSubscriber(logger *logrus.Logger) {
	var url = fmt.Sprintf("%s:%s", viper.GetString("atlas.pubsub.address"), viper.GetString("atlas.pubsub.port"))
	var topic = viper.GetString("atlas.pubsub.subscribe")
	var subscriptionID = viper.GetString("atlas.pubsub.subscriber.id")
	logger.Printf("pubsub: subscribing to server at %s with topic %q and subscription ID %q", url, topic, subscriptionID)
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("pubsub subscriber: Failed to dial to grpc server won't receive any messages %v", err)
	}
	s := pubsubgrpc.NewSubscriber(topic, subscriptionID, conn)
	c, e := s.Start(context.Background())
	for {
		select {
		case msg, isOpen := <-c:
			if !isOpen {
				logger.Println("pubsub: subscription channel closed")
				return
			}
			greeting := string(msg.Message())
			logger.Printf("pubsub: received message: %q", greeting)
			go func() {
				if err := msg.Ack(); err != nil {
					logger.Fatalf("pubsub: failed to ack messageID %q: %v", msg.MessageID(), err)
				}
			}()
		case err := <-e:
			logger.Printf("pubsub: encountered error reading subscription: %v", err)
		}
	}
}
