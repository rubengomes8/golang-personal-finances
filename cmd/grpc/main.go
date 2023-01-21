package main

import (
	"log"
	"net"

	"github.com/rubengomes8/golang-personal-finances/internal/enums"
	server "github.com/rubengomes8/golang-personal-finances/internal/grpc/expenses"
	"github.com/rubengomes8/golang-personal-finances/internal/pb/expenses"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/rds"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/rds/card"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/rds/expense"
	expenseRepo "github.com/rubengomes8/golang-personal-finances/internal/repository/rds/expense"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	database, err := rds.NewDB(
		enums.DatabaseHost,
		enums.DatabaseUser,
		enums.DatabasePwd,
		enums.DatabaseName,
		enums.DatabasePort,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	listener, err := net.Listen("tcp", enums.ListenerADDR)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	grpcServer := grpc.NewServer()

	cardRepo := card.NewCardRDS(database)
	expCategoryRepo := expense.NewCategoryRDS(database)
	expSubCategoryRepo := expense.NewSubCategoryRDS(database)
	expensesRepository := expenseRepo.NewRepo(database, cardRepo, expCategoryRepo, expSubCategoryRepo)

	expensesService, err := server.NewExpensesService(&expensesRepository, &expSubCategoryRepo, &cardRepo)
	if err != nil {
		log.Fatalf("Failed to create the finances server: %v\n", err)
	}

	expenses.RegisterExpensesServiceServer(grpcServer, &expensesService)

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
