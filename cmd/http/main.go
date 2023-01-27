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
	expensesRepo := database.NewExpensesRepo(db, cardRepo, expCategoryRepo, expSubCategoryRepo)

	incCategoryRepo := database.NewIncomeCategoryRepo(db)
	incomesRepo := database.NewIncomesRepo(db, cardRepo, incCategoryRepo)

	expensesService, err := service.NewExpenses(&expensesRepo, &expSubCategoryRepo, &cardRepo)
	if err != nil {
		log.Fatalf("Could not create expenses http service: %v\n", err)
	}

	incomesService, err := service.NewIncomes(&incomesRepo, &incCategoryRepo, &cardRepo)
	if err != nil {
		log.Fatalf("Could not create incomes http service: %v\n", err)
	}

	userRepo := database.NewUserRepo(db)
	authService, err := service.NewAuthService(&userRepo)
	if err != nil {
		log.Fatalf("Could not create auth http service: %v\n", err)
	}

	r := routes.SetupRouter(expensesService, incomesService, authService)
	err = r.Run()
	if err != nil {
		log.Fatalf("Could not run http router: %v\n", err)
	}
}
