package main

import (
	"log"

	"github.com/rubengomes8/golang-personal-finances/internal/enums"
	"github.com/rubengomes8/golang-personal-finances/internal/http/routes"
	"github.com/rubengomes8/golang-personal-finances/internal/http/service"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/rds"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/rds/card"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/rds/expense"

	_ "github.com/lib/pq"
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

	cardRepo := card.NewCardRDS(database)
	expCategoryRepo := expense.NewCategoryRDS(database)
	expSubCategoryRepo := expense.NewSubCategoryRDS(database)
	expensesRepository := expense.NewRepo(database, cardRepo, expCategoryRepo, expSubCategoryRepo)

	expensesControler, err := service.NewExpensesService(&expensesRepository, &expSubCategoryRepo, &cardRepo)
	if err != nil {
		log.Fatal(err)
	}

	r := routes.SetupRouter(expensesControler)
	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
