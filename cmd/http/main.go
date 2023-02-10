package main

import (
	"log"
	"os"
	"strconv"

	"github.com/rubengomes8/golang-personal-finances/internal/http/routes"
	"github.com/rubengomes8/golang-personal-finances/internal/http/service"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/database"

	_ "github.com/lib/pq"
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
