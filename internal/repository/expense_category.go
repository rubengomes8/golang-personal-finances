package repository

import (
	"context"

	"github.com/rubengomes8/golang-personal-finances/internal/models"
)

type ExpenseCategoryRepo interface {
	InsertExpenseCategory(context.Context, models.ExpenseCategory) (int64, error)
	UpdateExpenseCategory(context.Context, models.ExpenseCategory) (int64, error)
	GetExpenseCategoryByID(context.Context, int64) (models.ExpenseCategory, error)
	GetExpenseCategoryByName(context.Context, string) (models.ExpenseCategory, error)
	DeleteExpenseCategory(context.Context, int64) error
}
