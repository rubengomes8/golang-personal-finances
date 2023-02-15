package main

import (
	"log"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rubengomes8/golang-personal-finances/internal/http/handlers"
	"github.com/rubengomes8/golang-personal-finances/internal/http/routes"
	"github.com/rubengomes8/golang-personal-finances/internal/instrumentation"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/database/card"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/database/expense"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/database/income"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/database/user"
	service "github.com/rubengomes8/golang-personal-finances/internal/service/incomes"
	"github.com/rubengomes8/golang-personal-finances/internal/tools"

	_ "github.com/lib/pq"                                            //no lint
	_ "github.com/rubengomes8/golang-personal-finances/internal/env" //no lint
)

func main() {

	// INSTRUMENTATION
	instrumentation.Init()

	// DATABASE
	db, err := tools.InitPostgres(os.Getenv("DB_HOST"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	// REPOS
	prometheusLabels := prometheus.Labels{"version": "v1"}
	cardDB, err := card.NewCardRepoWithRED(
		card.NewDBWithLogs(card.NewDatabase(db)),
		prometheusLabels,
	)
	if err != nil {
		log.Fatalf("Failed to set up card repo with RED: %v\n", err)
	}

	expCategoryDB, err := expense.NewExpenseCategoryRepoWithRED(
		expense.NewExpenseCategoryRepoWithLogs(
			expense.NewCategoryDB(db)),
		prometheusLabels,
	)
	if err != nil {
		log.Fatalf("Failed to set up expense category repo with RED: %v\n", err)
	}

	expSubCategoryDB, err := expense.NewExpenseSubCategoryRepoWithRED(
		expense.NewExpenseSubCategoryRepoWithLogs(
			expense.NewSubCategoryDB(db)),
		prometheusLabels,
	)
	if err != nil {
		log.Fatalf("Failed to set up expense subcategory repo with RED: %v\n", err)
	}

	expensesDB, err := expense.NewExpenseRepoWithRED(
		expense.NewExpenseRepoWithLogs(
			expense.NewDB(db, cardDB, expCategoryDB, expSubCategoryDB),
		),
		prometheusLabels,
	)
	if err != nil {
		log.Fatalf("Failed to set up expense repo with RED: %v\n", err)
	}

	incCategoryDB, err := income.NewIncomeCategoryRepoWithRED(
		income.NewIncomeCategoryRepoWithLogs(income.NewCategoryDB(db)),
		prometheusLabels,
	)
	if err != nil {
		log.Fatalf("Failed to set up income category repo with RED: %v\n", err)
	}

	incomesDB, err := income.NewIncomeRepoWithRED(
		income.NewIncomeRepoWithLogs(income.NewDB(db, cardDB, incCategoryDB)),
		prometheusLabels,
	)
	if err != nil {
		log.Fatalf("Failed to set up incomes repo with RED: %v\n", err)
	}

	userDB, err := user.NewUserRepoWithRED(
		user.NewUserRepoWithLogs(user.NewDB(db)),
		prometheusLabels,
	)
	if err != nil {
		log.Fatalf("Failed to set up user repo with RED: %v\n", err)
	}

	// SERVICES
	// incomes service factory is using configuration pattern
	incomesService, err := service.NewIncomesWithConfiguration(
		service.WithIncomesRepository(incomesDB),
		service.WithCategoryRepository(incCategoryDB),
		service.WithCardRepository(cardDB),
	)
	if err != nil {
		log.Fatalf("Failed to set up incomes service with configuration patterns: %v\n", err)
	}

	// HTTP HANDLERS
	expensesHandlers := handlers.NewExpenses(expensesDB, expSubCategoryDB, cardDB)
	incomesHandlers := handlers.NewIncomes(incomesService)
	authHandlers := handlers.NewAuth(userDB)

	// HTTP ROUTER
	r := routes.SetupRouter(expensesHandlers, incomesHandlers, authHandlers)
	err = r.Run()
	if err != nil {
		log.Fatalf("Could not run http router: %v\n", err)
	}
}
