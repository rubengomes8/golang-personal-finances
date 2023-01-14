package repository

import (
	"context"
	"time"

	"github.com/rubengomes8/golang-personal-finances/internal/models"
)

type ExpenseRepo interface {
	InsertExpense(context.Context, models.Expense) (int64, error)
	UpdateExpense(context.Context, models.Expense) (int64, error)
	GetExpenseByID(context.Context, int64) (models.Expense, error)
	GetExpensesByDates(context.Context, time.Time, time.Time) ([]models.Expense, error)
	GetExpensesByCategory(context.Context, string) ([]models.Expense, error)
	GetExpensesBySubCategory(context.Context, string) ([]models.Expense, error)
	GetExpensesByCard(context.Context, string) ([]models.Expense, error)
	DeleteExpense(context.Context, int64) error
}
