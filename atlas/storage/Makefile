#include Makefile.vars
#include Makefile.common



# to override a target from Makefile.common just redefine the target.
# you can also chain the original atlas target by adding
# -atlas to the dependency of the redefined target

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

.PHONY: run-storage
run-storage:
	@dapr run --app-id storage \
           --app-protocol http \
           --app-port 8080 \
           --dapr-http-port 3500 \
           --log-level debug \
           --components-path ./config \
           make run-storage-server

.PHONY: run-storage-server
run-storage-server:
	@go run cmd/server/*.go \
			--database.address=localhost \
    		--database.port=5432  \
    		--database.user=postgres \
    		--database.password=postgres \
    		--database.ssl=disable \
    		--database.name=backend \
