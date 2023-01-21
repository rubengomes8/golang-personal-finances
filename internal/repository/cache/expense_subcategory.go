package cache

import (
	"context"

	models "github.com/rubengomes8/golang-personal-finances/internal/models/rds"
)

// ExpenseSubCategory implements the expense subCategory repository methods
type ExpenseSubCategory struct {
	repository []models.ExpenseSubCategoryTable
}

// NewExpenseSubCategory creates a ExpenseSubCategory cache
func NewExpenseSubCategory(repository []models.ExpenseSubCategoryTable) ExpenseSubCategory {
	return ExpenseSubCategory{
		repository: repository,
	}
}

// InsertExpenseSubCategory inserts an expense sub category on the cache if expense sub category does not exist
func (ecc *ExpenseSubCategory) InsertExpenseSubCategory(ctx context.Context, expSubCategory models.ExpenseSubCategoryTable) (int64, error) {

	existingCard, err := ecc.GetExpenseSubCategoryByID(ctx, expSubCategory.ID)
	if err == nil {
		return 0, SubCategoryAlreadyExistsError{
			id: existingCard.ID,
		}
	}

	ecc.repository = append(ecc.repository, expSubCategory)

	return 1, nil
}

// UpdateExpenseSubCategory updates an expense sub category on the cache if it exists
func (ecc *ExpenseSubCategory) UpdateExpenseSubCategory(ctx context.Context, updatedExpSubCategory models.ExpenseSubCategoryTable) (int64, error) {

	for idx, subCategory := range ecc.repository {
		if subCategory.ID == updatedExpSubCategory.ID {
			ecc.repository[idx] = updatedExpSubCategory
			return updatedExpSubCategory.ID, nil
		}
	}

	return 0, SubCategoryNotFoundByIDError{
		id: updatedExpSubCategory.ID,
	}
}

// GetExpenseSubCategoryByID returns the expense sub category from the cache if expense sub category with that id exists
func (ecc *ExpenseSubCategory) GetExpenseSubCategoryByID(ctx context.Context, id int64) (models.ExpenseSubCategoryTable, error) {

	for _, subCategory := range ecc.repository {
		if subCategory.ID == id {
			return subCategory, nil
		}
	}

	return models.ExpenseSubCategoryTable{}, SubCategoryNotFoundByIDError{
		id: id,
	}
}

// GetExpenseSubCategoryByName returns the expense sub category from the cache if expense sub category with that name exists
func (ecc *ExpenseSubCategory) GetExpenseSubCategoryByName(ctx context.Context, name string) (models.ExpenseSubCategoryTable, error) {

	for _, subCategory := range ecc.repository {
		if subCategory.Name == name {
			return subCategory, nil
		}
	}

	return models.ExpenseSubCategoryTable{}, SubCategoryNotFoundByNameError{
		name: name,
	}
}

// DeleteExpenseSubCategory deletes the expense sub category from cache if it exists
func (ecc *ExpenseSubCategory) DeleteExpenseSubCategory(ctx context.Context, id int64) error {

	for idx, subCategory := range ecc.repository {
		if subCategory.ID == id {
			ecc.repository = append(ecc.repository[:idx], ecc.repository[idx+1:]...)
			return nil
		}
	}

	return SubCategoryNotFoundByIDError{
		id: id,
	}
}
