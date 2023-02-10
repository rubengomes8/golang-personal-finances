package main

import (
	"log"

	"github.com/rubengomes8/golang-personal-finances/internal/http/handlers"
	"github.com/rubengomes8/golang-personal-finances/internal/http/routes"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/database"
	service "github.com/rubengomes8/golang-personal-finances/internal/service/incomes"
	"github.com/rubengomes8/golang-personal-finances/internal/tools"

	_ "github.com/lib/pq"
)

func main() {

	// DATABASE - TODO
	db, err := tools.InitPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	// REPOS
	cardRepo := database.NewCardRepo(db)

	expCategoryRepo := database.NewExpenseCategoryRepo(db)
	expSubCategoryRepo := database.NewExpenseSubCategoryRepo(db)
	expensesRepo := database.NewExpensesRepo(db, cardRepo, expCategoryRepo, expSubCategoryRepo)

	incCategoryRepo := database.NewIncomeCategoryRepo(db)
	incomesRepo := database.NewIncomesRepo(db, cardRepo, incCategoryRepo)

	userRepo := database.NewUserRepo(db)

	// SERVICES
	incomesService := service.NewIncomes(incomesRepo, incCategoryRepo, cardRepo)

	// HTTP HANDLERS
	expensesHandlers := handlers.NewExpenses(expensesRepo, expSubCategoryRepo, cardRepo)
	incomesHandlers := handlers.NewIncomes(incomesService)
	authHandlers := handlers.NewAuth(userRepo)

	r := routes.SetupRouter(expensesHandlers, incomesHandlers, authHandlers)
	err = r.Run()
	if err != nil {
		log.Fatalf("Could not run http router: %v\n", err)
	}
}
