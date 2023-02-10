package main

import (
	"log"
	"net"

	"github.com/rubengomes8/golang-personal-finances/internal/enums"
	grpcHandlers "github.com/rubengomes8/golang-personal-finances/internal/grpc"
	"github.com/rubengomes8/golang-personal-finances/internal/pb/expenses"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/database"
	"github.com/rubengomes8/golang-personal-finances/internal/tools"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {

	// DATABASE
	db, err := tools.InitPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	// REPOS
	cardRepo := database.NewCardRepo(db)
	expCategoryRepo := database.NewExpenseCategoryRepo(db)
	expSubCategoryRepo := database.NewExpenseSubCategoryRepo(db)
	expensesRepository := database.NewExpensesRepo(db, cardRepo, expCategoryRepo, expSubCategoryRepo)

	// HANDLERS / SERVICE
	expensesHandlers, err := grpcHandlers.NewExpenses(expensesRepository, expSubCategoryRepo, cardRepo)
	if err != nil {
		log.Fatalf("Failed to create the finances server: %v\n", err)
	}

	// TCP LISTERNER
	listener, err := net.Listen("tcp", enums.ListenerADDR)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	// GRPC SERVER
	grpcServer := grpc.NewServer()
	expenses.RegisterExpensesServiceServer(grpcServer, expensesHandlers)
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
