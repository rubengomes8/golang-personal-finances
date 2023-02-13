package main

import (
	"log"
	"os"

	"github.com/rubengomes8/golang-personal-finances/internal/http/handlers"
	"github.com/rubengomes8/golang-personal-finances/internal/http/routes"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/database"
	service "github.com/rubengomes8/golang-personal-finances/internal/service/incomes"
	"github.com/rubengomes8/golang-personal-finances/internal/tools"

	_ "github.com/lib/pq"                                            //no lint
	_ "github.com/rubengomes8/golang-personal-finances/internal/env" //no lint
)

func main() {

	// DATABASE
	db, err := tools.InitPostgres(os.Getenv("DB_HOST"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	// REPOS
	cardRepo := database.NewCardWithLogs(database.NewCard(db))

	expCategoryRepo := database.NewExpenseCategory(db)
	expSubCategoryRepo := database.NewExpenseSubCategory(db)
	expensesRepo := database.NewExpenses(db, cardRepo, expCategoryRepo, expSubCategoryRepo)

	incCategoryRepo := database.NewIncomeCategory(db)
	incomesRepo := database.NewIncomesWithLogs(database.NewIncomes(db, cardRepo, incCategoryRepo))

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
