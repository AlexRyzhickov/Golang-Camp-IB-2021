.PHONY: docker-build-generator
docker-build-generator:
	@docker build -t generator -f Dockerfile.generator .

.PHONE: gen-proto
gen-proto: docker-build-generator
	@docker run -d --rm \
		-v `pwd`/api:/api \
		-v `pwd`/internal/pb:/pb \
		generator protoc \
		--go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=. --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative --grpc-gateway_opt=generate_unbound_methods=true \
		-I /usr/local/include/. \
    	-I /api/. api.proto

.PHONY: run-db
run-db:
	@docker run \
		-d \
		--rm \
		-p 5432:5432 \
		--name db \
		-e POSTGRES_DB=backend \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=postgres \
		postgres:12

.PHONY: run-kube
run-kube:
	@kubectl apply -f manifest.yaml

.PHONY: docker-compose-up
docker-compose-up:
	@docker-compose up

.PHONY: gen-mocks
gen-mocks:
	docker run -v `pwd`:/src -w /src vektra/mockery:v2.9 --case snake --dir internal --output internal/mock --outpkg mock --all --exported

