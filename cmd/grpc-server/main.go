package main

import (
	"context"
	"log"
	"net"

	protoExpenses "github.com/rubengomes8/golang-personal-finances/proto/expenses"

	"google.golang.org/grpc"
)

const ADDR = "0.0.0.0:50051"

type FinancesServer struct {
	protoExpenses.ExpensesServiceServer
}

func main() {

	listener, err := net.Listen("tcp", ADDR)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", ADDR)

	grpcServer := grpc.NewServer()
	protoExpenses.RegisterExpensesServiceServer(grpcServer, &FinancesServer{})

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}

func (s *FinancesServer) CreateExpense(ctx context.Context, req *protoExpenses.ExpenseCreateRequest) (*protoExpenses.ExpenseCreateResponse, error) {
	log.Printf("CreateExpense was invoked with %v\n", req)
	return &protoExpenses.ExpenseCreateResponse{
		Id: 1,
	}, nil
}
