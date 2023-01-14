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
		return 0, fmt.Errorf("could not get card by name in insert expense: %v", err)
	}

	category, err := e.categoryRepo.GetExpenseCategoryByName(ctx, exp.Category)
	if err != nil {
		return 0, fmt.Errorf("could not get expense category by name in insert expense: %v", err)
	}

	subCategory, err := e.subCategoryRepo.GetExpenseSubCategoryByName(ctx, exp.SubCategory)
	if err != nil {
		return 0, fmt.Errorf("could not get expense subcategory by name in insert expense: %v", err)
	}

	insertStmt := fmt.Sprintf(`INSERT INTO %s 
	(value, date, description, category_id, subcategory_id, card_id) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`, tableName)

	var id int64
	err = e.database.QueryRowContext(ctx, insertStmt, exp.Value, ToTime(exp.Date), exp.Description, category.Id, subCategory.Id, card.Id).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not exec expense insert statement: %v", err)
	}

	return id, nil
}

/* UPDATE EXPENSE */
func (e *ExpenseRepo) UpdateExpense(ctx context.Context, exp models.Expense) (int64, error) {

	card, err := e.cardRepo.GetCardByName(ctx, exp.Card)
	if err != nil {
		return 0, fmt.Errorf("could not get card by name in update expense: %v", err)
	}

	category, err := e.categoryRepo.GetExpenseCategoryByName(ctx, exp.Category)
	if err != nil {
		return 0, fmt.Errorf("could not get expense category by name in update expense: %v", err)
	}

	subCategory, err := e.subCategoryRepo.GetExpenseSubCategoryByName(ctx, exp.SubCategory)
	if err != nil {
		return 0, fmt.Errorf("could not get expense subcategory by name in update expense: %v", err)
	}

	updateStmt := fmt.Sprintf(`UPDATE %s SET 
	(value, date, description, category_id, subcategory_id, card_id) =
	($1, $2, $3, $4, $5, $6) WHERE id = $7`, tableName)

	result, err := e.database.ExecContext(ctx, updateStmt, exp.Value, ToTime(exp.Date), exp.Description, category.Id, subCategory.Id, card.Id, exp.Id)
	if err != nil {
		return 0, fmt.Errorf("could not exec expense update statement: %v", err)
	}

	numRowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("could not get number of rows affected in exec expense update statement: %v", err)
	}

	if numRowsAffected == 0 {
		return 0, fmt.Errorf("there were no rows affected in exec expense update statement")
	}

	return exp.Id, nil
}

/* GET EXPENSE */
func (e *ExpenseRepo) GetExpenseByID(ctx context.Context, id int64) (models.ExpenseWithIDs, error) {

	selectStmt := fmt.Sprintf(`SELECT 
	(value, date, description, category_id, subcategory_id, card_id)
	FROM %s WHERE id = $1`, tableName)

	row := e.database.QueryRowContext(ctx, selectStmt, id)
	if row.Err() != nil {
		return models.ExpenseWithIDs{}, fmt.Errorf("could not query select by id expense statement: %v", row.Err())
	}

	var date time.Time
	exp := models.ExpenseWithIDs{Id: id}
	err := row.Scan(&exp.Value, &date, &exp.Description, &exp.CategoryId, &exp.SubCategoryId, &exp.CardId)
	if err != nil {
		return models.ExpenseWithIDs{}, fmt.Errorf("could not scan expense fields in get expense by id: %v", row.Err())
	}

	exp.Date = ToUnix(date)
	return exp, nil
}

/* GET EXPENSES */
func (e *ExpenseRepo) GetExpensesByDates(ctx context.Context, minDate int64, maxDate int64) ([]models.ExpenseWithIDs, error) {

	selectStmt := fmt.Sprintf(`SELECT 
	(value, date, description, category_id, subcategory_id, card_id) FROM %s 
	WHERE date BETWEEN $1 AND $2`, tableName)

	rows, err := e.database.QueryContext(ctx, selectStmt, minDate, maxDate)
	if err != nil {
		return []models.ExpenseWithIDs{}, fmt.Errorf("could not query select expenses by dates statement: %v", err)
	}

	var date time.Time
	var expenses []models.ExpenseWithIDs
	var exp models.ExpenseWithIDs
	for rows.Next() {
		err = rows.Scan(&exp.Value, &date, &exp.Description, &exp.CategoryId, &exp.SubCategoryId, &exp.CardId)
		if err != nil {
			return []models.ExpenseWithIDs{}, fmt.Errorf("could not scan expense fields in get expenses by dates: %v", err)
		}

		exp.Date = ToUnix(date)
		expenses = append(expenses, exp)
	}

	err = rows.Err()
	if err != nil {
		return []models.ExpenseWithIDs{}, fmt.Errorf("found error after scanning all expenses fields in get expenses by dates: %v", err)
	}

	return expenses, nil
}

func (e *ExpenseRepo) GetExpensesByCategory(ctx context.Context, category string) ([]models.ExpenseWithIDs, error) {

	expenseCategory, err := e.categoryRepo.GetExpenseCategoryByName(ctx, category)
	if err != nil {
		return []models.ExpenseWithIDs{}, fmt.Errorf("could not get expense category by name in get expenses by category: %v", err)
	}

	selectStmt := fmt.Sprintf(`SELECT 
	(value, date, description, category_id, subcategory_id, card_id) FROM %s 
	WHERE category_id = $1`, tableName)

	rows, err := e.database.QueryContext(ctx, selectStmt, expenseCategory.Id)
	if err != nil {
		return []models.ExpenseWithIDs{}, fmt.Errorf("could not query select expenses by category statement: %v", err)
	}

	var date time.Time
	var expenses []models.ExpenseWithIDs
	var exp models.ExpenseWithIDs
	for rows.Next() {
		err = rows.Scan(&exp.Value, &date, &exp.Description, &exp.CategoryId, &exp.SubCategoryId, &exp.CardId)
		if err != nil {
			return []models.ExpenseWithIDs{}, fmt.Errorf("could not scan expense fields in get expenses by category: %v", err)
		}

		exp.Date = ToUnix(date)
		expenses = append(expenses, exp)
	}

	err = rows.Err()
	if err != nil {
		return []models.ExpenseWithIDs{}, fmt.Errorf("found error after scanning all expenses fields in get expenses by category: %v", err)
	}

	return expenses, nil
}

func (e *ExpenseRepo) GetExpensesBySubCategory(ctx context.Context, subCategory string) ([]models.ExpenseWithIDs, error) {

	expenseSubCategory, err := e.subCategoryRepo.GetExpenseSubCategoryByName(ctx, subCategory)
	if err != nil {
		return []models.ExpenseWithIDs{}, fmt.Errorf("could not get expense subcategory by name in get expenses by subcategory: %v", err)
	}

	selectStmt := fmt.Sprintf(`SELECT 
	(value, date, description, category_id, subcategory_id, card_id) FROM %s 
	WHERE category_id = $1`, tableName)

	rows, err := e.database.QueryContext(ctx, selectStmt, expenseSubCategory.Id)
	if err != nil {
		return []models.ExpenseWithIDs{}, fmt.Errorf("could not query select expenses by subcategory statement: %v", err)
	}

	var date time.Time
	var expenses []models.ExpenseWithIDs
	var exp models.ExpenseWithIDs
	for rows.Next() {
		err = rows.Scan(&exp.Value, &date, &exp.Description, &exp.CategoryId, &exp.SubCategoryId, &exp.CardId)
		if err != nil {
			return []models.ExpenseWithIDs{}, fmt.Errorf("could not scan expense fields in get expenses by csubategory: %v", err)
		}

		exp.Date = ToUnix(date)
		expenses = append(expenses, exp)
	}

	err = rows.Err()
	if err != nil {
		return []models.ExpenseWithIDs{}, fmt.Errorf("found error after scanning all expenses fields in get expenses by subcategory: %v", err)
	}

	return expenses, nil
}

func (e *ExpenseRepo) GetExpensesByCard(ctx context.Context, card string) ([]models.ExpenseWithIDs, error) {
	return []models.ExpenseWithIDs{}, nil
}

/* DELETE */
func (e *ExpenseRepo) DeleteExpense(ctx context.Context, id int64) error {

	deleteStmt := fmt.Sprintf(`DELETE FROM %s 
	WHERE id = $1`, tableName)

	result, err := e.database.ExecContext(ctx, deleteStmt, id)
	if err != nil {
		return fmt.Errorf("could not exec expense delete statement: %v", err)
	}

	numRowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get number of rows affected in exec expense delete statement: %v", err)
	}

	if numRowsAffected == 0 {
		return fmt.Errorf("there were no rows affected in exec expense delete statement")
	}

	return nil

}

func ToTime(unixTime int64) time.Time {
	return time.Unix(unixTime, 0)
}

func ToUnix(time time.Time) int64 {
	return time.Unix()
}
