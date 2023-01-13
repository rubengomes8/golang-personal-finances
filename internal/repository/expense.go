package repository

import (
	"context"

	"github.com/rubengomes8/golang-personal-finances/internal/models"
)

type ExpenseRepo interface {
	InsertExpense(context.Context, models.Expense) (int64, error)
	UpdateExpense(context.Context, models.Expense) (int64, error)
	GetExpenseByID(context.Context, int64) (models.Expense, error)
	DeleteExpense(context.Context, int64) error
}
