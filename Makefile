GO_MODULE=github.com/rubengomes8/golang-personal-finances
BIN_DIR=bin

# DOCKER #
docker-up:
	docker-compose -f docker-compose.yaml up --build

docker-down: ## Stop docker containers and clear artefacts.
	docker-compose -f docker-compose.yaml down
	docker system prune

docker-image-prune: ## remove all images
	docker image prune -a --force


# PROTO GEN #
cards:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} cards.proto

expense_categories:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} expense_categories.proto

expense_subcategories:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} expense_subcategories.proto

expenses:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} expenses.proto

income_categories:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} income_categories.proto

incomes:
	protoc --proto_path=./proto --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} incomes.proto

all: cards expense_categories expense_subcategories expenses income_categories incomes


# BUILD #
build-grpc:
	go build -o ${BIN_DIR}/grpc_server ./cmd/grpc/main.go


# LINT #
lint:
	golangci-lint run --enable-all

# TEST #
test:
	go test ./...
