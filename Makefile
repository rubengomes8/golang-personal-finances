GO_MODULE=github.com/rubengomes8/golang-personal-finances
BIN_DIR=bin

# DOCKER #
up:
	docker-compose up -d --build

down:
	docker-compose down


# CARDS #
cards-create:
	protoc --proto_path=. --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} proto/cards/create.proto

cards-get:
	protoc --proto_path=. --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} proto/cards/get.proto

cards-service:
	protoc --proto_path=. --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} proto/cards/service.proto

cards: cards-create cards-get cards-service

# EXPENSES CATEGORIES #
expense_categories-create:
	protoc --proto_path=. --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} proto/expense_categories/create.proto

expense_categories-get:
	protoc --proto_path=. --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} proto/expense_categories/get.proto

expense_categories-service:
	protoc --proto_path=. --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} proto/expense_categories/service.proto

expense_categories: expense_categories-create expense_categories-get expense_categories-service

# EXPENSES SUB CATEGORIES #
expense_subcategories-create:
	protoc --proto_path=. --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} proto/expense_subcategories/create.proto

expense_subcategories-get:
	protoc --proto_path=. --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} proto/expense_subcategories/get.proto

expense_subcategories-service:
	protoc --proto_path=. --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} proto/expense_subcategories/service.proto

expense_subcategories: expense_subcategories-create expense_subcategories-get expense_subcategories-service

# EXPENSES #
expenses-create:
	protoc --proto_path=. --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} proto/expenses/create.proto

expenses-get:
	protoc --proto_path=. --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} proto/expenses/get.proto

expenses-service:
	protoc --proto_path=. --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} proto/expenses/service.proto

expenses: expenses-create expenses-get expenses-service

# ALL
all: cards expense_categories expense_subcategories expenses 

# BUILD GO #
build-expenses:
	go build -o ${BIN_DIR}/expenses/server ./cmd/grpc-server/main.go
	go build -o ${BIN_DIR}/expenses/client ./cmd/grpc-client/main.go

build: build-expenses