GO_MODULE=github.com/rubengomes8/golang-personal-finances

# DOCKER #
up:
	docker-compose up -d --build

down:
	docker-compose down


# EXPENSES #
expenses-create:
	protoc -Iproto/expenses --go_out=. --go_opt=module=${GO_MODULE} --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} proto/expenses/create.proto

expenses: expenses-create