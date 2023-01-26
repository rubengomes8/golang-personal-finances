FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /http_server ./cmd/http/main.go

EXPOSE 8080

CMD [ "/http_server"]