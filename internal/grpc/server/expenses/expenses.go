package server

import (
	"database/sql"

	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	"github.com/rubengomes8/golang-personal-finances/proto/expenses"
)

type ExpensesService struct {
	expenses.ExpensesServiceServer
	ExpensesRepository            repository.ExpenseRepo
	ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
	CardRepository                repository.CardRepo
	Database                      *sql.DB
}

func NewExpensesService(
	expRepo repository.ExpenseRepo,
	expSubCatRepo repository.ExpenseSubCategoryRepo,
	cardRepo repository.CardRepo,
	database *sql.DB,
) (ExpensesService, error) {
	return ExpensesService{
		ExpensesRepository:            expRepo,
		ExpensesSubCategoryRepository: expSubCatRepo,
		CardRepository:                cardRepo,
		Database:                      database,
	}, nil
}
