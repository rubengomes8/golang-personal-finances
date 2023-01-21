package expense

import (
	"context"

	models "github.com/rubengomes8/golang-personal-finances/internal/models/rds"
)

// ExpenseCategoryCache implements the expense category repository methods
type ExpenseCategoryCache struct {
	repository []models.ExpenseCategoryTable
}

// InsertExpenseCategory inserts an expense category on the cache if expense category does not exist
func (ecc *ExpenseCategoryCache) InsertExpenseCategory(ctx context.Context, expCategory models.ExpenseCategoryTable) (int64, error) {

	existingCard, err := ecc.GetExpenseCategoryByID(ctx, expCategory.ID)
	if err == nil {
		return 0, CategoryAlreadyExistsError{
			id: existingCard.ID,
		}
	}

	ecc.repository = append(ecc.repository, expCategory)

	return 1, nil
}

// UpdateExpenseCategory updates an expense category on the cache if it exists
func (ecc *ExpenseCategoryCache) UpdateExpenseCategory(ctx context.Context, updatedExpCategory models.ExpenseCategoryTable) (int64, error) {

	for idx, category := range ecc.repository {
		if category.ID == updatedExpCategory.ID {
			ecc.repository[idx] = updatedExpCategory
			return updatedExpCategory.ID, nil
		}
	}

	return 0, CategoryNotFoundByIDError{
		id: updatedExpCategory.ID,
	}
}

// GetExpenseCategoryByID returns the expense category from the cache if expense category with that id exists
func (ecc *ExpenseCategoryCache) GetExpenseCategoryByID(ctx context.Context, id int64) (models.ExpenseCategoryTable, error) {

	for _, category := range ecc.repository {
		if category.ID == id {
			return category, nil
		}
	}

	return models.ExpenseCategoryTable{}, CategoryNotFoundByIDError{
		id: id,
	}
}

// GetExpenseCategoryByName returns the expense category from the cache if expense category with that name exists
func (ecc *ExpenseCategoryCache) GetExpenseCategoryByName(ctx context.Context, name string) (models.ExpenseCategoryTable, error) {

	for _, category := range ecc.repository {
		if category.Name == name {
			return category, nil
		}
	}

	return models.ExpenseCategoryTable{}, CategoryNotFoundByNameError{
		name: name,
	}
}

// DeleteExpenseCategory deletes the expense category from cache if it exists
func (ecc *ExpenseCategoryCache) DeleteExpenseCategory(ctx context.Context, id int64) error {

	for idx, category := range ecc.repository {
		if category.ID == id {
			ecc.repository = append(ecc.repository[:idx], ecc.repository[idx+1:]...)
			return nil
		}
	}

	return CategoryNotFoundByIDError{
		id: id,
	}
}
