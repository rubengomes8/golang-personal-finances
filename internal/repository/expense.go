package repository

import (
	"context"
	"time"

	models "github.com/rubengomes8/golang-personal-finances/internal/models/rds"
)

// ExpenseRepo defines the expense repository interface
type ExpenseRepo interface {
	InsertExpense(context.Context, models.ExpenseTable) (int64, error)
	UpdateExpense(context.Context, models.ExpenseTable) (int64, error)
	GetExpenseByID(context.Context, int64) (models.ExpenseView, error)
	GetExpensesByDates(context.Context, time.Time, time.Time) ([]models.ExpenseView, error)
	GetExpensesByCategory(context.Context, string) ([]models.ExpenseView, error)
	GetExpensesBySubCategory(context.Context, string) ([]models.ExpenseView, error)
	GetExpensesByCard(context.Context, string) ([]models.ExpenseView, error)
	DeleteExpense(context.Context, int64) error
}
