package mock

import (
	"context"
	"errors"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

const (
	IncomeSalaryCategoryName = "Salary"
)

var (
	IncomeSalaryCategory = models.IncomeCategoryTable{
		ID:   1,
		Name: IncomeSalaryCategoryName,
	}
)

// IncomeCategory mocks the income category repository methods
type IncomeCategory struct {
}

// NewIncomeCategory creates a IncomeCategory mock
func NewIncomeCategory() IncomeCategory {
	return IncomeCategory{}
}

// InsertIncomeCategory mocks an income category insert
func (ic IncomeCategory) InsertIncomeCategory(ctx context.Context, income models.IncomeCategoryTable) (int64, error) {

	switch income.Name {
	case IncomeSalaryCategoryName:
		return IncomeSalaryCategory.ID, nil
	default:
		return 0, errors.New("could not insert income category")
	}
}

// UpdateIncomeCategory mocks an income category update
func (ic IncomeCategory) UpdateIncomeCategory(ctx context.Context, income models.IncomeCategoryTable) (int64, error) {

	switch income.Name {
	case IncomeSalaryCategoryName:
		return IncomeSalaryCategory.ID, nil
	default:
		return 0, errors.New("income category with this name does not exist")
	}
}

// UpdateIncomeCategory mocks an income category update
func (ic IncomeCategory) GetIncomeCategoryByID(ctx context.Context, id int64) (models.IncomeCategoryTable, error) {

	switch id {
	case IncomeSalaryCategory.ID:
		return IncomeSalaryCategory, nil
	default:
		return models.IncomeCategoryTable{}, errors.New("income category with this id does not exist")
	}
}

// UpdateIncomeCategory mocks a get income category by name
func (ic IncomeCategory) GetIncomeCategoryByName(ctx context.Context, name string) (models.IncomeCategoryTable, error) {
	switch name {
	case IncomeSalaryCategoryName:
		return IncomeSalaryCategory, nil
	default:
		return models.IncomeCategoryTable{}, errors.New("income category with this name does not exist")
	}
}

// DeleteIncomeCategory mocks a delete income
func (ic IncomeCategory) DeleteIncomeCategory(ctx context.Context, id int64) error {
	switch id {
	case IncomeSalaryCategory.ID:
		return nil
	default:
		return errors.New("income category with this id does not exist")
	}
}
