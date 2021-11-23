package main

import (
	models "atlas/storage/internal/model"
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/infobloxopen/atlas-app-toolkit/server"

	"github.com/infobloxopen/atlas-app-toolkit/gorm/resource"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

func main() {
	doneC := make(chan error)
	logger := NewLogger()
	if viper.GetBool("internal.enable") {
		go func() { doneC <- ServeInternal(logger) }()
	}

	go func() { doneC <- ServeExternal(logger) }()

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

var sub = &common.Subscription{
	PubsubName: "messages",
	Topic:      "neworder",
	Route:      "/orders",
}

// ServeExternal builds and runs the server that listens on ServerAddress and GatewayAddress
func ServeExternal(logger *logrus.Logger) error {

	if viper.GetString("database.dsn") == "" {
		setDBConnection()
	}

	db, err := gorm.Open(postgres.Open(viper.GetString("database.dsn")), &gorm.Config{})
	if err != nil {
		logger.Fatalln(err)
	}

	if isInit := db.Migrator().HasTable(&models.Note{}); !isInit {
		err := db.Migrator().CreateTable(&models.Note{})
		if err != nil {
			logger.Fatalln(err)
		}
	}

	s := daprd.NewService(":8080")

	if err := s.AddTopicEventHandler(sub, eventHandler); err != nil {
		logger.Fatalf("error adding topic subscription: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("error listenning: %v", err)
	}

	return nil
}

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)
	return false, nil
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

func forwardResponseOption(ctx context.Context, w http.ResponseWriter, resp protoreflect.ProtoMessage) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
	return nil
}

// setDBConnection sets the db connection string
func setDBConnection() {
	viper.Set("database.dsn", fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s dbname=%s",
		viper.GetString("database.address"), viper.GetString("database.port"),
		viper.GetString("database.user"), viper.GetString("database.password"),
		viper.GetString("database.ssl"), viper.GetString("database.name")))
}
