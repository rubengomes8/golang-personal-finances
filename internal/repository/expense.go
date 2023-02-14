package repository

import (
	"context"
	"time"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

//go:generate gowrap gen -g -i ExpenseRepo -t ./templates/log_template.go.tmpl -o ./database/expense/with_logs_by_template.go
//go:generate gowrap gen -g -i ExpenseRepo -t ./templates/red_template.go.tmpl -o ./database/expense/with_red_by_template.go
// ExpenseRepo defines the expense repository interface.
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
