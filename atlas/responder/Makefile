#include Makefile.vars
#include Makefile.common



# to override a target from Makefile.common just redefine the target.
# you can also chain the original atlas target by adding
# -atlas to the dependency of the redefined target

.PHONY: docker-build-generator
docker-build-generator:
	@docker build -t generator -f Dockerfile.generator .

.PHONE: gen-proto
gen-proto: docker-build-generator
	@docker run -d --rm \
		-v `pwd`/api:/api \
		-v `pwd`/pkg/pb:/pb \
		generator protoc \
		--go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=. --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative --grpc-gateway_opt=generate_unbound_methods=true \
		-I /usr/local/include/. \
    	-I /api/. api.proto

.PHONY: run-responder
run-responder:
	@export DAPR_PUBSUB_NAME=messages && dapr run --app-id responder \
           --app-protocol http \
           --app-port 8088 \
           --dapr-http-port 3501 \
           --log-level debug \
           --components-path ./config \
           go run cmd/server/*.go

.PHONY: run-dapr
run-dapr:
	@dapr run --app-id responder \
           --app-protocol http \
           --app-port 8088 \
           --dapr-http-port 3501 \
           --dapr-grpc-port 43011 \
           --log-level debug \
           --components-path ./config \

