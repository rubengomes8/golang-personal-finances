package repository

import (
	"context"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

//go:generate gowrap gen -g -i ExpenseSubCategoryRepo -t ./templates/log_template.go.tmpl -o ./database/expense/subcategory_with_logs_by_template.go
//go:generate gowrap gen -g -i ExpenseSubCategoryRepo -t ./templates/red_template.go.tmpl -o ./database/expense/subcategory_with_red_by_template.go
// ExpenseSubCategoryRepo defines the expense subcategory repository interface
type ExpenseSubCategoryRepo interface {
	InsertExpenseSubCategory(context.Context, models.ExpenseSubCategoryTable) (int64, error)
	UpdateExpenseSubCategory(context.Context, models.ExpenseSubCategoryTable) (int64, error)
	GetExpenseSubCategoryByID(context.Context, int64) (models.ExpenseSubCategoryTable, error)
	GetExpenseSubCategoryByName(context.Context, string) (models.ExpenseSubCategoryTable, error)
	DeleteExpenseSubCategory(context.Context, int64) error
}
