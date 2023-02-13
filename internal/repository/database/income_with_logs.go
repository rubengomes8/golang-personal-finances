package database

import (
	"context"
	"log"
	"time"

	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

// IncomesWithLogs is the Incomes decorator adding the logs
type IncomesWithLogs struct {
	repo repository.IncomeRepo
}

func NewIncomesWithLogs(repo Incomes) IncomesWithLogs {
	return IncomesWithLogs{
		repo: repo,
	}
}

func (i IncomesWithLogs) InsertIncome(ctx context.Context, income models.IncomeTable) (int64, error) {
	log.Printf("income: %+v", income)
	return i.repo.InsertIncome(ctx, income)
}
func (i IncomesWithLogs) UpdateIncome(ctx context.Context, income models.IncomeTable) (int64, error) {
	log.Printf("income: %+v", income)
	return i.repo.UpdateIncome(ctx, income)
}

func (i IncomesWithLogs) GetIncomeByID(ctx context.Context, id int64) (models.IncomeView, error) {
	log.Printf("income id: %+v", id)
	return i.repo.GetIncomeByID(ctx, id)
}

func (i IncomesWithLogs) GetIncomesByDates(ctx context.Context, minDate time.Time, maxDate time.Time) ([]models.IncomeView, error) {
	log.Printf("income dates: min_date: %+v | max_date: %+v", minDate, maxDate)
	return i.repo.GetIncomesByDates(ctx, minDate, maxDate)
}

func (i IncomesWithLogs) GetIncomesByCategory(ctx context.Context, category string) ([]models.IncomeView, error) {
	log.Printf("income category: %+v", category)
	return i.repo.GetIncomesByCategory(ctx, category)
}

func (i IncomesWithLogs) GetIncomesByCard(ctx context.Context, card string) ([]models.IncomeView, error) {
	log.Printf("income card: %+v", card)
	return i.repo.GetIncomesByCard(ctx, card)
}

func (i IncomesWithLogs) DeleteIncome(ctx context.Context, id int64) error {
	log.Printf("income id: %+v", id)
	return i.repo.DeleteIncome(ctx, id)
}
