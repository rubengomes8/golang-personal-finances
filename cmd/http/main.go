package main

import (
	"log"

	"github.com/rubengomes8/golang-personal-finances/internal/enums"
	"github.com/rubengomes8/golang-personal-finances/internal/http/controllers"
	"github.com/rubengomes8/golang-personal-finances/internal/http/routes"
	"github.com/rubengomes8/golang-personal-finances/internal/postgres"
	"github.com/rubengomes8/golang-personal-finances/internal/postgres/card"
	"github.com/rubengomes8/golang-personal-finances/internal/postgres/expense"

	_ "github.com/lib/pq"
)

func main() {

	database, err := postgres.NewDB(
		enums.DatabaseHost,
		enums.DatabaseUser,
		enums.DatabasePwd,
		enums.DatabaseName,
		enums.DatabasePort,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	cardRepo := card.NewCardRepo(database)
	expCategoryRepo := expense.NewCategoryRepo(database)
	expSubCategoryRepo := expense.NewSubCategoryRepo(database)
	expensesRepository := expense.NewRepo(database, cardRepo, expCategoryRepo, expSubCategoryRepo)

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
