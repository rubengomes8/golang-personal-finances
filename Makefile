GO_MODULE=github.com/rubengomes8/golang-personal-finances
BIN_DIR=bin

# DOCKER #
up:
	docker-compose up -d --build

down:
	docker-compose down


# EXPENSES #
expenses-create:
	protoc --proto_path=. --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} proto/expenses/create.proto

expenses-get:
	protoc --proto_path=. --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} proto/expenses/get.proto

expenses-service:
	protoc --proto_path=. --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} proto/expenses/service.proto

expenses: expenses-create expenses-get expenses-service

# BUILD GO #
build-expenses:
	go build -o ${BIN_DIR}/expenses/server ./cmd/grpc-server/main.go
	go build -o ${BIN_DIR}/expenses/client ./cmd/grpc-client/main.go

build: build-expenses