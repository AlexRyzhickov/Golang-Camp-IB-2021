module atlas/storage

go 1.16

replace github.com/spf13/afero => github.com/spf13/afero v1.5.1

require (
	github.com/dapr/go-sdk v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0 // indirect
	github.com/infobloxopen/atlas-app-toolkit v1.1.1
	github.com/prometheus/client_golang v1.11.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.9.0
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/postgres v1.2.2
	gorm.io/gorm v1.22.3
)
