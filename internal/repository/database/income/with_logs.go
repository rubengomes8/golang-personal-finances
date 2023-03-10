package income

import (
	"context"
	"log"
	"time"

	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

// DBWithLogs is the Incomes decorator adding the logs
type DBWithLogs struct {
	repo repository.IncomeRepo
}

func NewDBWithLogs(repo DB) DBWithLogs {
	return DBWithLogs{
		repo: repo,
	}
}

func (i DBWithLogs) InsertIncome(ctx context.Context, income models.IncomeTable) (int64, error) {
	log.Printf("income: %+v", income)
	return i.repo.InsertIncome(ctx, income)
}
func (i DBWithLogs) UpdateIncome(ctx context.Context, income models.IncomeTable) (int64, error) {
	log.Printf("income: %+v", income)
	return i.repo.UpdateIncome(ctx, income)
}

func (i DBWithLogs) GetIncomeByID(ctx context.Context, id int64) (models.IncomeView, error) {
	log.Printf("income id: %+v", id)
	return i.repo.GetIncomeByID(ctx, id)
}

func (i DBWithLogs) GetIncomesByDates(ctx context.Context, minDate time.Time, maxDate time.Time) ([]models.IncomeView, error) {
	log.Printf("income dates: min_date: %+v | max_date: %+v", minDate, maxDate)
	return i.repo.GetIncomesByDates(ctx, minDate, maxDate)
}

func (i DBWithLogs) GetIncomesByCategory(ctx context.Context, category string) ([]models.IncomeView, error) {
	log.Printf("income category: %+v", category)
	return i.repo.GetIncomesByCategory(ctx, category)
}

func (i DBWithLogs) GetIncomesByCard(ctx context.Context, card string) ([]models.IncomeView, error) {
	log.Printf("income card: %+v", card)
	return i.repo.GetIncomesByCard(ctx, card)
}

func (i DBWithLogs) DeleteIncome(ctx context.Context, id int64) error {
	log.Printf("income id: %+v", id)
	return i.repo.DeleteIncome(ctx, id)
}
