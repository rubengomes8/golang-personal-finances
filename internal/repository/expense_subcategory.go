package repository

import (
	"context"

	models "github.com/rubengomes8/golang-personal-finances/internal/models/rds"
)

type ExpenseSubCategoryRepo interface {
	InsertExpenseSubCategory(context.Context, models.ExpenseSubCategoryTable) (int64, error)
	UpdateExpenseSubCategory(context.Context, models.ExpenseSubCategoryTable) (int64, error)
	GetExpenseSubCategoryByID(context.Context, int64) (models.ExpenseSubCategoryTable, error)
	GetExpenseSubCategoryByName(context.Context, string) (models.ExpenseSubCategoryTable, error)
	DeleteExpenseSubCategory(context.Context, int64) error
}
