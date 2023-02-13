package repository

import (
	"context"
	"time"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

// Incomes defines the expense repository interface.
type Incomes interface {
	InsertIncome(context.Context, models.IncomeTable) (int64, error)
	UpdateIncome(context.Context, models.IncomeTable) (int64, error)
	GetIncomeByID(context.Context, int64) (models.IncomeView, error)
	GetIncomesByDates(context.Context, time.Time, time.Time) ([]models.IncomeView, error)
	GetIncomesByCategory(context.Context, string) ([]models.IncomeView, error)
	GetIncomesByCard(context.Context, string) ([]models.IncomeView, error)
	DeleteIncome(context.Context, int64) error
}
