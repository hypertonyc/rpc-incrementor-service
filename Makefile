PROD_POSTGRES_CONTAINER_NAME=postgres-container
PROD_POSTGRES_DBNAME=incrementor
PROD_POSTGRES_PASSWORD=super_secret_123
PROD_POSTGRES_DATA_PATH=./data
PROD_GRPC_CONTAINER_NAME=incrementor-container
PROD_GRPC_DB_CONNECTION=postgres://postgres:super_secret_123@postgresql/incrementor?sslmode=disable
PROD_GRPC_PORT=9000

TEST_POSTGRES_CONTAINER_NAME=postgres-container-test
TEST_POSTGRES_DBNAME=incrementor-test
TEST_POSTGRES_PASSWORD=secret
TEST_POSTGRES_DATA_PATH=./data-test
TEST_GRPC_CONTAINER_NAME=incrementor-container-test
TEST_GRPC_DB_CONNECTION=postgres://postgres:secret@postgresql/incrementor-test?sslmode=disable
TEST_GRPC_PORT=9009
TEST_RESULT_FILE=result.log
TEST_SUCCESS_RESULT_FILE=result-success.log

PROTOBUF_CODEGEN_DIR=internal/api/grpc/pb

prod-env:
	export CURRENT_UID=$$(id -u) && \
	export CURRENT_GID=$$(id -g) && \
	export POSTGRES_CONTAINER_NAME=$(PROD_POSTGRES_CONTAINER_NAME) && \
	export POSTGRES_DB=$(PROD_POSTGRES_DBNAME) && \
	export POSTGRES_PASSWORD=$(PROD_POSTGRES_PASSWORD) && \
	export POSTGRES_DATA_PATH=$(PROD_POSTGRES_DATA_PATH) && \
	export GRPC_CONTAINER_NAME=$(PROD_GRPC_CONTAINER_NAME) && \
	export GRPC_DB_CONNECTION=$(PROD_GRPC_DB_CONNECTION) && \
	export GRPC_PORT=$(PROD_GRPC_PORT) && \
	$(CMD)

test-env:
	export CURRENT_UID=$$(id -u) && \
	export CURRENT_GID=$$(id -g) && \
	export POSTGRES_CONTAINER_NAME=$(TEST_POSTGRES_CONTAINER_NAME) && \
	export POSTGRES_DB=$(TEST_POSTGRES_DBNAME) && \
	export POSTGRES_PASSWORD=$(TEST_POSTGRES_PASSWORD) && \
	export POSTGRES_DATA_PATH=$(TEST_POSTGRES_DATA_PATH) && \
	export GRPC_CONTAINER_NAME=$(TEST_GRPC_CONTAINER_NAME) && \
	export GRPC_DB_CONNECTION=$(TEST_GRPC_DB_CONNECTION) && \
	export GRPC_PORT=$(TEST_GRPC_PORT) && \
	$(CMD)

proto:
	rm -f $(PROTOBUF_CODEGEN_DIR)/*.go
	protoc --proto_path internal/api/grpc/proto \
		--go_out=$(PROTOBUF_CODEGEN_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(PROTOBUF_CODEGEN_DIR) --go-grpc_opt=paths=source_relative \
		internal/api/grpc/proto/*.proto

test:
	go test -v -short ./internal/domain
	go test -v -short ./internal/repository
	go test -v -short ./internal/service

debug:
	go run cmd/incrementor/main.go --inmemory

e2e: CMD=rm -f $(TEST_RESULT_FILE) && \
rm -fR $$POSTGRES_DATA_PATH && \
mkdir -p $$POSTGRES_DATA_PATH && \
docker compose up -d && \
make e2e-scenario || true && \
diff -c $(TEST_SUCCESS_RESULT_FILE) $(TEST_RESULT_FILE) || true && \
docker compose stop && \
docker compose rm -f && \
rm -fR $$POSTGRES_DATA_PATH && \
rm -f $(TEST_RESULT_FILE)
e2e: test-env

e2e-scenario:
	sleep 5
	docker run --net incrementor_network --rm incrementor /app/incrementor-cli get_number >> $(TEST_RESULT_FILE)
	docker run --net incrementor_network --rm incrementor /app/incrementor-cli increment
	docker run --net incrementor_network --rm incrementor /app/incrementor-cli increment
	docker run --net incrementor_network --rm incrementor /app/incrementor-cli get_number >> $(TEST_RESULT_FILE)
	docker run --net incrementor_network --rm incrementor /app/incrementor-cli set_settings --limit 20 --step 7
	docker run --net incrementor_network --rm incrementor /app/incrementor-cli increment
	docker run --net incrementor_network --rm incrementor /app/incrementor-cli get_number >> $(TEST_RESULT_FILE)
	docker run --net incrementor_network --rm incrementor /app/incrementor-cli increment
	docker run --net incrementor_network --rm incrementor /app/incrementor-cli get_number >> $(TEST_RESULT_FILE)
	docker run --net incrementor_network --rm incrementor /app/incrementor-cli set_settings --limit 20 --step 30 >> $(TEST_RESULT_FILE) 2>&1 || true
	docker run --net incrementor_network --rm incrementor /app/incrementor-cli set_settings --limit 0 --step 30 >> $(TEST_RESULT_FILE) 2>&1 || true
	docker run --net incrementor_network --rm incrementor /app/incrementor-cli set_settings --limit 20 --step 0 >> $(TEST_RESULT_FILE) 2>&1 || true
	docker compose restart grpc-server
	sleep 5
	docker run --net incrementor_network --rm incrementor /app/incrementor-cli get_number >> $(TEST_RESULT_FILE)
	docker run --net incrementor_network --rm incrementor /app/incrementor-cli increment
	docker run --net incrementor_network --rm incrementor /app/incrementor-cli get_number >> $(TEST_RESULT_FILE)

up: CMD=mkdir -p $$POSTGRES_DATA_PATH && docker compose up -d
up: prod-env

stop: CMD=docker compose stop
stop: prod-env

.PHONY: prod-env test-env proto test debug e2e e2e-scenario up stop