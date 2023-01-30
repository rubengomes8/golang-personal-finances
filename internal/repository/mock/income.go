package mock

import (
	"context"
	"errors"
	"time"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

var (
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
		Category:    IncomeSalaryCategory.Name,
		Card:        IncomeSalaryCard.Name,
		CategoryID:  IncomeSalaryCategory.ID,
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

	switch income.Description {
	case "Mock":
		return 1, nil
	default:
		return 0, errors.New("could not insert income")
	}
}

// InsertIncome mocks an income update
func (i Income) UpdateIncome(ctx context.Context, income models.IncomeTable) (int64, error) {

	switch income.ID {
	case 1:
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

// GetIncomesByDates mocks an income get by dates
func (i Income) GetIncomesByDates(ctx context.Context, min time.Time, max time.Time) ([]models.IncomeView, error) {

	if min.Before(IncomeSalaryDate) && max.After(IncomeSalaryDate) {
		return []models.IncomeView{
			IncomeSalaryView,
		}, nil
	}

	return []models.IncomeView{}, errors.New("could not get income view by dates")
}

// GetIncomesByCategory mocks an income get by category
func (i Income) GetIncomesByCategory(ctx context.Context, category string) ([]models.IncomeView, error) {

	if category == IncomeSalaryCategory.Name {
		return []models.IncomeView{
			IncomeSalaryView,
		}, nil
	}

	return []models.IncomeView{}, errors.New("could not get income view by category")
}

// GetIncomesByCard mocks an income get by card
func (i Income) GetIncomesByCard(ctx context.Context, card string) ([]models.IncomeView, error) {

	if card == IncomeSalaryCard.Name {
		return []models.IncomeView{
			IncomeSalaryView,
		}, nil
	}

	return []models.IncomeView{}, errors.New("could not get income view by card")
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
