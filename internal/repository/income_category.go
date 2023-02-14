package repository

import (
	"context"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

//go:generate gowrap gen -g -i IncomeCategoryRepo -t ./templates/log_template.go.tmpl -o ./database/income/category_with_logs_by_template.go
//go:generate gowrap gen -g -i IncomeCategoryRepo -t ./templates/red_template.go.tmpl -o ./database/income/category_with_red_by_template.go
// IncomeCategoryRepo defines the income category repository interface.
type IncomeCategoryRepo interface {
	InsertIncomeCategory(context.Context, models.IncomeCategoryTable) (int64, error)
	UpdateIncomeCategory(context.Context, models.IncomeCategoryTable) (int64, error)
	GetIncomeCategoryByID(context.Context, int64) (models.IncomeCategoryTable, error)
	GetIncomeCategoryByName(context.Context, string) (models.IncomeCategoryTable, error)
	DeleteIncomeCategory(context.Context, int64) error
}
