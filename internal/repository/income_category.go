package repository

import (
	"context"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

// IncomeCategoryRepo defines the income category repository interface.
type IncomeCategoryRepo interface {
	InsertIncomeCategory(context.Context, models.IncomeCategoryTable) (int64, error)
	UpdateIncomeCategory(context.Context, models.IncomeCategoryTable) (int64, error)
	GetIncomeCategoryByID(context.Context, int64) (models.ExpenseCategoryTable, error)
	GetIncomeCategoryByName(context.Context, string) (models.ExpenseCategoryTable, error)
	DeleteIncomeCategory(context.Context, int64) error
}
