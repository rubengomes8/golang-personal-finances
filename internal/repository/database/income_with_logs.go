package database

import (
	"context"
	"time"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

type IncomesWithLogs struct {
	repo Incomes
}

func NewIncomesWithLogs(repo Incomes) IncomesWithLogs {
	return IncomesWithLogs{
		repo: repo,
	}
}

func (i IncomesWithLogs) InsertIncome(ctx context.Context, income models.IncomeTable) (int64, error) {
	return i.repo.InsertIncome(ctx, income)
}
func (i IncomesWithLogs) UpdateIncome(ctx context.Context, income models.IncomeTable) (int64, error) {
	return i.repo.UpdateIncome(ctx, income)
}

func (i IncomesWithLogs) GetIncomeByID(ctx context.Context, id int64) (models.IncomeView, error) {
	return i.repo.GetIncomeByID(ctx, id)
}

func (i IncomesWithLogs) GetIncomesByDates(ctx context.Context, minDate time.Time, maxDate time.Time) ([]models.IncomeView, error) {
	return i.repo.GetIncomesByDates(ctx, minDate, maxDate)
}

func (i IncomesWithLogs) GetIncomesByCategory(ctx context.Context, category string) ([]models.IncomeView, error) {
	return i.repo.GetIncomesByCategory(ctx, category)
}

func (i IncomesWithLogs) GetIncomesByCard(ctx context.Context, card string) ([]models.IncomeView, error) {
	return i.repo.GetIncomesByCard(ctx, card)
}

func (i IncomesWithLogs) DeleteIncome(ctx context.Context, id int64) error {
	return i.repo.DeleteIncome(ctx, id)
}
