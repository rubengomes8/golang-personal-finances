package server

import (
	"database/sql"

	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	"github.com/rubengomes8/golang-personal-finances/proto/expenses"
)

type FinancesServer struct {
	expenses.ExpensesServiceServer
	ExpensesRepository repository.ExpenseRepo
	Database           *sql.DB
}

func New(expensesRepository repository.ExpenseRepo, database *sql.DB) (FinancesServer, error) {
	return FinancesServer{
		ExpensesRepository: expensesRepository,
		Database:           database,
	}, nil
}
