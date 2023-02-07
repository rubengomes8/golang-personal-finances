package main

import (
	"log"

	"github.com/rubengomes8/golang-personal-finances/internal/enums"
	"github.com/rubengomes8/golang-personal-finances/internal/http/handlers"
	"github.com/rubengomes8/golang-personal-finances/internal/http/routes"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/database"
	"github.com/rubengomes8/golang-personal-finances/internal/services/incomes"

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

	// Repos
	cardRepo := database.NewCardRepo(db)

	expCategoryRepo := database.NewExpenseCategoryRepo(db)
	expSubCategoryRepo := database.NewExpenseSubCategoryRepo(db)
	expensesRepo := database.NewExpensesRepo(db, cardRepo, expCategoryRepo, expSubCategoryRepo)

	incCategoryRepo := database.NewIncomeCategoryRepo(db)
	incomesRepo := database.NewIncomesRepo(db, cardRepo, incCategoryRepo)

	userRepo := database.NewUserRepo(db)

	// Services
	incomesService := incomes.New(&incomesRepo, &incCategoryRepo, &cardRepo)

	// HTTP Handlers
	expensesHandlers, err := handlers.NewExpenses(&expensesRepo, &expSubCategoryRepo, &cardRepo)
	if err != nil {
		log.Fatalf("Could not create expenses http service: %v\n", err)
	}

	incomesHandlers, err := handlers.NewIncomes(incomesService)
	if err != nil {
		log.Fatalf("Could not create incomes http service: %v\n", err)
	}

	authHandlers, err := handlers.NewAuthService(&userRepo)
	if err != nil {
		log.Fatalf("Could not create auth http service: %v\n", err)
	}

	// HTTP Router
	r := routes.SetupRouter(expensesHandlers, incomesHandlers, authHandlers)
	err = r.Run()
	if err != nil {
		log.Fatalf("Could not run http router: %v\n", err)
	}
}
