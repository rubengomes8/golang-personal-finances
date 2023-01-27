package main

import (
	"log"

	"github.com/rubengomes8/golang-personal-finances/internal/enums"
	"github.com/rubengomes8/golang-personal-finances/internal/http/routes"
	"github.com/rubengomes8/golang-personal-finances/internal/http/service"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/database"

	_ "github.com/lib/pq"
)

func main() {

	db, err := database.New(
		enums.DatabaseHost,
		enums.DatabaseUser,
		enums.DatabasePwd,
		enums.DatabaseName,
		enums.DatabasePort,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	cardRepo := database.NewCardRepo(db)
	expCategoryRepo := database.NewExpenseCategoryRepo(db)
	expSubCategoryRepo := database.NewExpenseSubCategoryRepo(db)
	expensesRepository := database.NewExpensesRepo(db, cardRepo, expCategoryRepo, expSubCategoryRepo)

	expensesService, err := service.NewExpenses(&expensesRepository, &expSubCategoryRepo, &cardRepo)
	if err != nil {
		log.Fatalf("Could not create expenses http service: %v\n", err)
	}

	userRepo := database.NewUserRepo(db)
	authService, err := service.NewAuthService(&userRepo)
	if err != nil {
		log.Fatalf("Could not create auth http service: %v\n", err)
	}

	r := routes.SetupRouter(expensesService, authService)
	err = r.Run()
	if err != nil {
		log.Fatalf("Could not run http router: %v\n", err)
	}
}
