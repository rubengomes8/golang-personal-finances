package repository

import (
	"context"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

//go:generate gowrap gen -g -i ExpenseCategoryRepo -t ./templates/log_template.go.tmpl -o ./database/expense/category_with_logs_by_template.go
//go:generate gowrap gen -g -i ExpenseCategoryRepo -t ./templates/red_template.go.tmpl -o ./database/expense/category_with_red_by_template.go
// ExpenseCategoryRepo defines the expense category repository interface.
type ExpenseCategoryRepo interface {
	InsertExpenseCategory(context.Context, models.ExpenseCategoryTable) (int64, error)
	UpdateExpenseCategory(context.Context, models.ExpenseCategoryTable) (int64, error)
	GetExpenseCategoryByID(context.Context, int64) (models.ExpenseCategoryTable, error)
	GetExpenseCategoryByName(context.Context, string) (models.ExpenseCategoryTable, error)
	DeleteExpenseCategory(context.Context, int64) error
}
