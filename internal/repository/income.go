package repository

import (
	"context"
	"time"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

//go:generate gowrap gen -g -i IncomeRepo -t ./templates/log_template.go.tmpl -o ./database/income/with_logs_by_template.go
//go:generate gowrap gen -g -i IncomeRepo -t ./templates/red_template.go.tmpl -o ./database/income/with_red_by_template.go
// IncomeRepo defines the incomes repository interface.
type IncomeRepo interface {
	InsertIncome(context.Context, models.IncomeTable) (int64, error)
	UpdateIncome(context.Context, models.IncomeTable) (int64, error)
	GetIncomeByID(context.Context, int64) (models.IncomeView, error)
	GetIncomesByDates(context.Context, time.Time, time.Time) ([]models.IncomeView, error)
	GetIncomesByCategory(context.Context, string) ([]models.IncomeView, error)
	GetIncomesByCard(context.Context, string) ([]models.IncomeView, error)
	DeleteIncome(context.Context, int64) error
}
