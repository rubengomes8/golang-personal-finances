package main

import (
	"log"
	"net"

	"github.com/rubengomes8/golang-personal-finances/internal/grpc/server"
	"github.com/rubengomes8/golang-personal-finances/proto/expenses"

	"google.golang.org/grpc"
)

const ADDR = "0.0.0.0:50051"

func main() {

	listener, err := net.Listen("tcp", ADDR)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", ADDR)

	grpcServer := grpc.NewServer()
	expenses.RegisterExpensesServiceServer(grpcServer, &server.FinancesServer{})

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
