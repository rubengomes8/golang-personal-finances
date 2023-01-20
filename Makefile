GO_MODULE=github.com/rubengomes8/golang-personal-finances
BIN_DIR=bin

# DOCKER #
up:
	docker-compose up -d --build

down:
	docker-compose down

# CARDS #
cards:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} cards.proto

# EXPENSES CATEGORIES #
expense_categories:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} expense_categories.proto

# EXPENSES SUB CATEGORIES #
expense_subcategories:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} expense_subcategories.proto

# EXPENSES #
expenses:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} expenses.proto

all: cards expense_categories expense_subcategories expenses

# BUILD GO #
build-expenses:
	go build -o ${BIN_DIR}/expenses/server ./cmd/grpc-server/main.go

build: build-expenses

# LINT #
lint:
	golangci-lint run --enable-all

# TEST #
test:
	go test ./...
