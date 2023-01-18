package main

import (
	"log"

	"github.com/rubengomes8/golang-personal-finances/internal/http/controllers"
	"github.com/rubengomes8/golang-personal-finances/internal/http/routes"
	"github.com/rubengomes8/golang-personal-finances/internal/postgres"
	"github.com/rubengomes8/golang-personal-finances/internal/postgres/card"
	"github.com/rubengomes8/golang-personal-finances/internal/postgres/expense"
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

	cardRepo := card.NewCardRepo(database)
	expCategoryRepo := expense.NewExpenseCategoryRepo(database)
	expSubCategoryRepo := expense.NewExpenseSubCategoryRepo(database)
	expensesRepository := expense.NewExpenseRepo(database, cardRepo, expCategoryRepo, expSubCategoryRepo)

	expensesControler, err := controllers.NewExpensesController(&expensesRepository, &expSubCategoryRepo, &cardRepo)
	if err != nil {
		log.Fatal(err)
	}

	r := routes.SetupRouter(expensesControler)
	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
