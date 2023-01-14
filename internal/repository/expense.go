package repository

import (
	"context"

	"github.com/rubengomes8/golang-personal-finances/internal/models"
)

type ExpenseRepo interface {
	InsertExpense(context.Context, models.Expense) (int64, error)
	UpdateExpense(context.Context, models.Expense) (int64, error)
	GetExpenseByID(context.Context, int64) (models.ExpenseWithIDs, error)
	GetExpensesByDates(context.Context, int64, int64) ([]models.ExpenseWithIDs, error)
	GetExpensesByCategory(context.Context, string) ([]models.ExpenseWithIDs, error)
	GetExpensesBySubCategory(context.Context, string) ([]models.ExpenseWithIDs, error)
	GetExpensesByCard(context.Context, string) ([]models.ExpenseWithIDs, error)
	DeleteExpense(context.Context, int64) error
}
