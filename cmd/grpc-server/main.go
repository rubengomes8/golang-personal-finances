package main

import (
	"log"
	"net"

	server "github.com/rubengomes8/golang-personal-finances/internal/grpc/server/expenses"
	"github.com/rubengomes8/golang-personal-finances/internal/pb/expenses"
	"github.com/rubengomes8/golang-personal-finances/internal/postgres"
	"github.com/rubengomes8/golang-personal-finances/internal/postgres/card"
	"github.com/rubengomes8/golang-personal-finances/internal/postgres/expense"
	expenseRepo "github.com/rubengomes8/golang-personal-finances/internal/postgres/expense"

	"google.golang.org/grpc"
)

const (
	LISTENER_ADDR = "0.0.0.0:50051"

	DATABASE_HOST = "0.0.0.0"
	DATABASE_PORT = 5432
	DATABASE_USER = "finances@ruben"
	DATABASE_PWD  = "rub3nF!n4nc3s"
	DATABASE_NAME = "finances"
)

func main() {

	database, err := postgres.NewDB(DATABASE_HOST, DATABASE_USER, DATABASE_PWD, DATABASE_NAME, DATABASE_PORT)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	listener, err := net.Listen("tcp", LISTENER_ADDR)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on %s\n", LISTENER_ADDR)

	grpcServer := grpc.NewServer()

	cardRepo := card.NewCardRepo(database)
	expCategoryRepo := expense.NewExpenseCategoryRepo(database)
	expSubCategoryRepo := expense.NewExpenseSubCategoryRepo(database)

	expensesRepository := expenseRepo.NewExpenseRepo(database, cardRepo, expCategoryRepo, expSubCategoryRepo)
	expensesService, err := server.NewExpensesService(&expensesRepository, &expSubCategoryRepo, &cardRepo, database)
	if err != nil {
		log.Fatalf("Failed to create the finances server: %v\n", err)
	}
	expenses.RegisterExpensesServiceServer(grpcServer, &expensesService)

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
