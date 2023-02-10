package main

import (
	"log"
	"net"
	"os"
	"strconv"

	"github.com/rubengomes8/golang-personal-finances/internal/enums"
	"github.com/rubengomes8/golang-personal-finances/internal/grpc/service"
	"github.com/rubengomes8/golang-personal-finances/internal/pb/expenses"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/database"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {

	dbPortEnv := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(dbPortEnv)
	if err != nil {
		log.Fatalf("Could not convert database port to interger: %v\n", err)
	}

	db, err := database.New(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_NAME"),
		dbPort,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	listener, err := net.Listen("tcp", enums.ListenerADDR)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	grpcServer := grpc.NewServer()

	cardRepo := database.NewCardRepo(db)
	expCategoryRepo := database.NewExpenseCategoryRepo(db)
	expSubCategoryRepo := database.NewExpenseSubCategoryRepo(db)
	expensesRepository := database.NewExpensesRepo(db, cardRepo, expCategoryRepo, expSubCategoryRepo)

	expensesService, err := service.NewExpenses(&expensesRepository, &expSubCategoryRepo, &cardRepo)
	if err != nil {
		log.Fatalf("Failed to create the finances server: %v\n", err)
	}

	expenses.RegisterExpensesServiceServer(grpcServer, &expensesService)

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
