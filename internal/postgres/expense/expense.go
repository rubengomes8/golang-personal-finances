package expense

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rubengomes8/golang-personal-finances/internal/models"
	"github.com/rubengomes8/golang-personal-finances/internal/postgres/card"
)

const (
	tableName = "expenses"
)

type ExpenseRepo struct {
	database        *sql.DB
	cardRepo        card.CardRepo
	categoryRepo    ExpenseCategoryRepo
	subCategoryRepo ExpenseSubCategoryRepo
}

func NewExpenseRepo(
	database *sql.DB,
	cardRepo card.CardRepo,
	categoryRepo ExpenseCategoryRepo,
	subCategoryRepo ExpenseSubCategoryRepo,
) ExpenseRepo {
	return ExpenseRepo{
		database:        database,
		cardRepo:        cardRepo,
		categoryRepo:    categoryRepo,
		subCategoryRepo: subCategoryRepo,
	}
}

/* INSERT EXPENSE */
func (e *ExpenseRepo) InsertExpense(ctx context.Context, exp models.Expense) (int64, error) {

	card, err := e.cardRepo.GetCardByName(ctx, exp.Card)
	if err != nil {
		return 0, fmt.Errorf("could not get card by name: %v", err)
	}

	category, err := e.categoryRepo.GetExpenseCategoryByName(ctx, exp.Category)
	if err != nil {
		return 0, fmt.Errorf("could not get expense category by name: %v", err)
	}

	subCategory, err := e.subCategoryRepo.GetExpenseSubCategoryByName(ctx, exp.SubCategory)
	if err != nil {
		return 0, fmt.Errorf("could not get expense subcategory by name: %v", err)
	}

	insertStmt := fmt.Sprintf(`INSERT INTO %s 
	(value, date, description, category_id, subcategory_id, card_id) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`, tableName)

	var id int64
	err = e.database.QueryRowContext(ctx, insertStmt, exp.Value, ToTime(exp.Date), exp.Description, category.Id, subCategory.Id, card.Id).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not exec insert expense statement: %v", err)
	}

	return id, nil
}

/* UPDATE EXPENSE */
func (e *ExpenseRepo) UpdateExpense(ctx context.Context, exp models.Expense) (int64, error) {

	card, err := e.cardRepo.GetCardByName(ctx, exp.Card)
	if err != nil {
		return 0, fmt.Errorf("could not get card by name: %v", err)
	}

	category, err := e.categoryRepo.GetExpenseCategoryByName(ctx, exp.Category)
	if err != nil {
		return 0, fmt.Errorf("could not get expense category by name: %v", err)
	}

	subCategory, err := e.subCategoryRepo.GetExpenseSubCategoryByName(ctx, exp.SubCategory)
	if err != nil {
		return 0, fmt.Errorf("could not get expense subcategory by name: %v", err)
	}

	updateStmt := fmt.Sprintf(`UPDATE %s SET 
	(value, date, description, category_id, subcategory_id, card_id) =
	($1, $2, $3, $4, $5, $6) WHERE id = $7`, tableName)

	result, err := e.database.ExecContext(ctx, updateStmt, exp.Value, ToTime(exp.Date), exp.Description, category.Id, subCategory.Id, card.Id, exp.Id)
	if err != nil {
		return 0, fmt.Errorf("could not exec update expense statement: %v", err)
	}

	numRowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("could not get number of rows affected in exec update statement: %v", err)
	}

	if numRowsAffected == 0 {
		return 0, fmt.Errorf("there were no rows affected in exec update statement")
	}

	return exp.Id, nil
}

/* GET EXPENSE */
func (e *ExpenseRepo) GetExpenseByID(ctx context.Context, id int64) (models.Expense, error) {
	return models.Expense{}, nil
}

/* GET EXPENSES */
func (e *ExpenseRepo) GetExpensesByDates(ctx context.Context, minDate time.Time, maxDate time.Time) ([]models.Expense, error) {
	return []models.Expense{}, nil
}

func (e *ExpenseRepo) GetExpensesByCategory(ctx context.Context, category string) ([]models.Expense, error) {
	return []models.Expense{}, nil
}

func (e *ExpenseRepo) GetExpensesBySubCategory(ctx context.Context, subCategory string) ([]models.Expense, error) {
	return []models.Expense{}, nil
}

func (e *ExpenseRepo) GetExpensesByCard(ctx context.Context, card string) ([]models.Expense, error) {
	return []models.Expense{}, nil
}

/* DELETE */
func (e *ExpenseRepo) DeleteExpense(ctx context.Context, id int64) error {
	return nil
}

func ToTime(unixTime int64) time.Time {
	return time.Unix(unixTime, 0)
}
