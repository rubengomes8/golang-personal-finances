package mock

import (
	"context"
	"errors"
	"time"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

var (
	IncomeSalaryCard = models.CardTable{
		ID:   1,
		Name: "CGD",
	}

	IncomeSalaryDate = time.Now().UTC()
	IncomeSalary     = models.IncomeTable{
		ID:          1,
		Value:       1000,
		Date:        IncomeSalaryDate,
		CategoryID:  1,
		CardID:      IncomeSalaryCard.ID,
		Description: "Mock",
	}

	IncomeSalaryView = models.IncomeView{
		ID:          IncomeSalary.ID,
		Value:       IncomeSalary.Value,
		Date:        IncomeSalary.Date,
		Category:    IncomeCategorySalary.Name,
		Card:        IncomeSalaryCard.Name,
		CategoryID:  IncomeCategorySalary.ID,
		CardID:      IncomeSalaryCard.ID,
		Description: IncomeSalary.Description,
	}
)

// Income mocks the income repository methods
type Income struct {
}

// NewIncome creates a Income mock
func NewIncome() Income {
	return Income{}
}

// InsertIncome mocks an income insert
func (i Income) InsertIncome(ctx context.Context, income models.IncomeTable) (int64, error) {
	switch income {
	case IncomeSalary:
		return 1, nil
	default:
		return 0, errors.New("could not insert income")
	}
}

// InsertIncome mocks an income update
func (i Income) UpdateIncome(ctx context.Context, income models.IncomeTable) (int64, error) {
	switch income {
	case IncomeSalary:
		return 1, nil
	default:
		return 0, errors.New("could not update income")
	}
}

// InsertIncome mocks an income get by id
func (i Income) GetIncomeByID(ctx context.Context, id int64) (models.IncomeView, error) {
	switch id {
	case IncomeSalary.ID:
		return IncomeSalaryView, nil
	default:
		return models.IncomeView{}, errors.New("could not get income view by id")
	}
}

// GetIncomesByDates - TODO
func (i Income) GetIncomesByDates(context.Context, time.Time, time.Time) ([]models.IncomeView, error) {
	return []models.IncomeView{}, nil
}

// GetIncomesByCategory - TODO
func (i Income) GetIncomesByCategory(context.Context, string) ([]models.IncomeView, error) {
	return []models.IncomeView{}, nil

}

// GetIncomesByCard - TODO
func (i Income) GetIncomesByCard(context.Context, string) ([]models.IncomeView, error) {
	return []models.IncomeView{}, nil

}

// DeleteIncome mocks an income delete
func (i Income) DeleteIncome(ctx context.Context, id int64) error {

	switch id {
	case IncomeSalary.ID:
		return nil
	default:
		return errors.New("income with this id does not exist")
	}
}
