FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /grpc_server ./cmd/grpc/main.go

EXPOSE 50051

CMD [ "/grpc_server"]