package repository

import (
	"context"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

// ExpenseCategoryRepo defines the expense category repository interface.
type ExpenseCategoryRepo interface {
	InsertExpenseCategory(context.Context, models.ExpenseCategoryTable) (int64, error)
	UpdateExpenseCategory(context.Context, models.ExpenseCategoryTable) (int64, error)
	GetExpenseCategoryByID(context.Context, int64) (models.ExpenseCategoryTable, error)
	GetExpenseCategoryByName(context.Context, string) (models.ExpenseCategoryTable, error)
	DeleteExpenseCategory(context.Context, int64) error
}
