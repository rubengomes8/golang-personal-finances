GO_MODULE=github.com/rubengomes8/golang-personal-finances
BIN_DIR=bin

# DOCKER #
up:
	docker-compose up -d --build

down:
	docker-compose down

# PROTO GEN #
cards:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} cards.proto

expense_categories:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} expense_categories.proto

expense_subcategories:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} expense_subcategories.proto

expenses:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} expenses.proto

all: cards expense_categories expense_subcategories expenses


# BUILD GO #
build-expenses:
	go build -o ${BIN_DIR}/grpc_server ./cmd/grpc/main.go
	go build -o ${BIN_DIR}/http_server ./cmd/http/main.go

build: build-expenses

# LINT #
lint:
	golangci-lint run --enable-all

# TEST #
test:
	go test ./...
