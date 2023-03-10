package cache

import (
	"context"
	"time"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

// Expense implements the expense repository methods
type Expense struct {
	cardrepository        Card
	categoryrepository    ExpenseCategory
	subCategoryrepository ExpenseSubCategory
	repository            []models.ExpenseTable
}

// NewExpense creates a Expense cache
func NewExpense(
	repository []models.ExpenseTable,
	cardRepo Card,
	catRepo ExpenseCategory,
	subCatRepo ExpenseSubCategory,
) Expense {
	return Expense{
		cardrepository:        cardRepo,
		categoryrepository:    catRepo,
		subCategoryrepository: subCatRepo,
		repository:            repository,
	}
}

// InsertExpense inserts an expense on the cache if expense category does not exist
func (ec *Expense) InsertExpense(ctx context.Context, e models.ExpenseTable) (int64, error) {

	ec.repository = append(ec.repository, e)

	return 1, nil
}

// UpdateExpense updates an expense on the cache if it exists
func (ec *Expense) UpdateExpense(ctx context.Context, e models.ExpenseTable) (int64, error) {

	for idx, exp := range ec.repository {
		if exp.ID == e.ID {
			ec.repository[idx] = e
			return e.ID, nil
		}
	}

	return 0, ExpenseNotFoundByIDError{
		id: e.ID,
	}
}

// GetExpenseByID returns the expense from the cache if expense with that id exists
func (ec *Expense) GetExpenseByID(ctx context.Context, id int64) (models.ExpenseView, error) {

	var expense models.ExpenseTable
	for _, exp := range ec.repository {
		if exp.ID == id {
			expense = exp
			break
		}
	}

	cardTable, err := ec.cardrepository.GetCardByID(ctx, expense.CardID)
	if err != nil {
		return models.ExpenseView{}, GettingCardByIDError{
			id: expense.CardID,
		}
	}

	subCategoryTable, err := ec.subCategoryrepository.GetExpenseSubCategoryByID(ctx, expense.SubCategoryID)
	if err != nil {
		return models.ExpenseView{}, GettingSubCategoryByIDError{
			id: expense.SubCategoryID,
		}
	}

	categoryTable, err := ec.categoryrepository.GetExpenseCategoryByID(ctx, subCategoryTable.CategoryID)
	if err != nil {
		return models.ExpenseView{}, GettingCategoryByIDError{
			id: subCategoryTable.CategoryID,
		}
	}

	return models.ExpenseView{
		ID:            expense.ID,
		Value:         expense.Value,
		Date:          expense.Date,
		Category:      categoryTable.Name,
		SubCategory:   subCategoryTable.Name,
		Card:          cardTable.Name,
		CategoryID:    categoryTable.ID,
		SubCategoryID: expense.SubCategoryID,
		CardID:        expense.CardID,
		Description:   expense.Description,
	}, nil
}

// GetExpensesByDates returns the expenses from the cache if expense with that dates' range exists
func (ec *Expense) GetExpensesByDates(ctx context.Context, minDate time.Time, maxDate time.Time) ([]models.ExpenseView, error) {

	var expenseViews []models.ExpenseView
	for _, exp := range ec.repository {

		if exp.Date.After(minDate) && exp.Date.Before(maxDate) {

			subCategoryTable, err := ec.subCategoryrepository.GetExpenseSubCategoryByID(ctx, exp.SubCategoryID)
			if err != nil {
				return []models.ExpenseView{}, GettingSubCategoryByIDError{
					id: exp.SubCategoryID,
				}
			}

			cardTable, err := ec.cardrepository.GetCardByID(ctx, exp.CardID)
			if err != nil {
				return []models.ExpenseView{}, GettingCardByIDError{
					id: exp.CardID,
				}
			}

			categoryTable, err := ec.categoryrepository.GetExpenseCategoryByID(ctx, subCategoryTable.CategoryID)
			if err != nil {
				return []models.ExpenseView{}, GettingCategoryByIDError{
					id: subCategoryTable.CategoryID,
				}
			}

			expenseViews = append(expenseViews,
				models.ExpenseView{
					ID:            exp.ID,
					Value:         exp.Value,
					Date:          exp.Date,
					Category:      categoryTable.Name,
					SubCategory:   subCategoryTable.Name,
					Card:          cardTable.Name,
					CategoryID:    categoryTable.ID,
					SubCategoryID: subCategoryTable.ID,
					CardID:        cardTable.ID,
					Description:   exp.Description,
				})
		}
	}

	return expenseViews, nil
}

// GetExpensesByCategory returns the expenses from the cache if expense with that category exists
func (ec *Expense) GetExpensesByCategory(ctx context.Context, cat string) ([]models.ExpenseView, error) {

	categoryTable, err := ec.categoryrepository.GetExpenseCategoryByName(ctx, cat)
	if err != nil {
		return []models.ExpenseView{}, GettingCategoryByNameError{
			name: cat,
		}
	}

	var expenseViews []models.ExpenseView
	for _, exp := range ec.repository {

		subCategoryTable, err := ec.subCategoryrepository.GetExpenseSubCategoryByID(ctx, exp.SubCategoryID)
		if err != nil {
			return []models.ExpenseView{}, GettingSubCategoryByIDError{
				id: exp.SubCategoryID,
			}
		}

		if categoryTable.ID == subCategoryTable.CategoryID {

			cardTable, err := ec.cardrepository.GetCardByID(ctx, exp.CardID)
			if err != nil {
				return []models.ExpenseView{}, GettingCardByIDError{
					id: exp.CardID,
				}
			}

			expenseViews = append(expenseViews,
				models.ExpenseView{
					ID:            exp.ID,
					Value:         exp.Value,
					Date:          exp.Date,
					Category:      categoryTable.Name,
					SubCategory:   subCategoryTable.Name,
					Card:          cardTable.Name,
					CategoryID:    categoryTable.ID,
					SubCategoryID: subCategoryTable.ID,
					CardID:        cardTable.ID,
					Description:   exp.Description,
				})
		}
	}

	return expenseViews, nil
}

// GetExpensesBySubCategory returns the expenses from the cache if expense with that subcategory exists
func (ec *Expense) GetExpensesBySubCategory(ctx context.Context, subCat string) ([]models.ExpenseView, error) {

	subCategoryTable, err := ec.subCategoryrepository.GetExpenseSubCategoryByName(ctx, subCat)
	if err != nil {
		return []models.ExpenseView{}, GettingSubCategoryByNameError{
			name: subCat,
		}
	}

	categoryTable, err := ec.categoryrepository.GetExpenseCategoryByID(ctx, subCategoryTable.CategoryID)
	if err != nil {
		return []models.ExpenseView{}, GettingCategoryByIDError{
			id: subCategoryTable.CategoryID,
		}
	}

	var expenseViews []models.ExpenseView
	for _, exp := range ec.repository {
		if exp.SubCategoryID == subCategoryTable.ID {

			cardTable, err := ec.cardrepository.GetCardByID(ctx, exp.CardID)
			if err != nil {
				return []models.ExpenseView{}, GettingCardByIDError{
					id: exp.CardID,
				}
			}

			expenseViews = append(expenseViews,
				models.ExpenseView{
					ID:            exp.ID,
					Value:         exp.Value,
					Date:          exp.Date,
					Category:      categoryTable.Name,
					SubCategory:   subCategoryTable.Name,
					Card:          cardTable.Name,
					CategoryID:    categoryTable.ID,
					SubCategoryID: subCategoryTable.ID,
					CardID:        cardTable.ID,
					Description:   exp.Description,
				})
		}
	}

	return expenseViews, nil
}

// GetExpensesByCard returns the expenses from the cache if expense with that card exists
func (ec *Expense) GetExpensesByCard(ctx context.Context, card string) ([]models.ExpenseView, error) {

	cardTable, err := ec.cardrepository.GetCardByName(ctx, card)
	if err != nil {
		return []models.ExpenseView{}, GettingCardByNameError{
			name: card,
		}
	}

	var expenseViews []models.ExpenseView
	for _, exp := range ec.repository {
		if cardTable.ID == exp.CardID {

			subCategoryTable, err := ec.subCategoryrepository.GetExpenseSubCategoryByID(ctx, exp.SubCategoryID)
			if err != nil {
				return []models.ExpenseView{}, GettingSubCategoryByIDError{
					id: exp.SubCategoryID,
				}
			}

			categoryTable, err := ec.categoryrepository.GetExpenseCategoryByID(ctx, subCategoryTable.CategoryID)
			if err != nil {
				return []models.ExpenseView{}, GettingCategoryByIDError{
					id: subCategoryTable.CategoryID,
				}
			}

			expenseViews = append(expenseViews,
				models.ExpenseView{
					ID:            exp.ID,
					Value:         exp.Value,
					Date:          exp.Date,
					Category:      categoryTable.Name,
					SubCategory:   subCategoryTable.Name,
					Card:          cardTable.Name,
					CategoryID:    categoryTable.ID,
					SubCategoryID: subCategoryTable.ID,
					CardID:        cardTable.ID,
					Description:   exp.Description,
				})
		}
	}

	return expenseViews, nil
}

// DeleteExpense deletes the expense from cache if it exists
func (ec *Expense) DeleteExpense(ctx context.Context, id int64) error {

	for idx, expense := range ec.repository {
		if expense.ID == id {
			ec.repository = append(ec.repository[:idx], ec.repository[idx+1:]...)
			return nil
		}
	}

	return ExpenseNotFoundByIDError{
		id: id,
	}
}
