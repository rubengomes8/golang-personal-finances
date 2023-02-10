package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

const (
	expensesTable = "expenses"
	expensesView  = "expenses_view"
)

// ExpensesRepo implements the expense repository methods
type ExpensesRepo struct {
	database     *sql.DB
	cardRepo     CardRepo
	categoryRepo ExpenseCategoryRepo
	subCategory  ExpenseSubCategoryRepo
}

// NewExpensesRepo creates a new ExpensesRepo
func NewExpensesRepo(
	database *sql.DB,
	cardRepo CardRepo,
	categoryRepo ExpenseCategoryRepo,
	subCategory ExpenseSubCategoryRepo,
) ExpensesRepo {
	return ExpensesRepo{
		database:     database,
		cardRepo:     cardRepo,
		categoryRepo: categoryRepo,
		subCategory:  subCategory,
	}
}

// InsertExpense inserts an expense on the expenses db table
func (e ExpensesRepo) InsertExpense(ctx context.Context, exp models.ExpenseTable) (int64, error) {

	insertStmt := fmt.Sprintf(`INSERT INTO %s 
	(value, date, description, subcategory_id, card_id)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`, expensesTable)

	var id int64

	err := e.database.QueryRowContext(
		ctx,
		insertStmt,
		exp.Value,
		exp.Date,
		exp.Description,
		exp.SubCategoryID,
		exp.CardID,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not exec expense insert statement: %v", err)
	}

	return id, nil
}

// UpdateExpense updates an expense on the expenses db table
func (e ExpensesRepo) UpdateExpense(ctx context.Context, exp models.ExpenseTable) (int64, error) {

	updateStmt := fmt.Sprintf(`UPDATE %s SET 
	(value, date, description, subcategory_id, card_id) =
	($1, $2, $3, $4, $5) WHERE id = $6`, expensesTable)

	result, err := e.database.ExecContext(ctx,
		updateStmt,
		exp.Value,
		exp.Date,
		exp.Description,
		exp.SubCategoryID,
		exp.CardID,
		exp.ID,
	)
	if err != nil {
		return 0, fmt.Errorf("could not exec expense update statement: %v", err)
	}

	numRowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("could not get number of rows affected in exec expense update statement: %v", err)
	}

	if numRowsAffected == 0 {
		return 0, ErrNoRowsAffectedExpenseUpdate
	}

	return exp.ID, nil
}

// GetExpenseByID gets an expense from the expenses db table by id
func (e ExpensesRepo) GetExpenseByID(ctx context.Context, id int64) (models.ExpenseView, error) {

	selectStmt := fmt.Sprintf(`SELECT 
	value, date, description, category_id, category_name, 
	subcategory_id, subcategory_name, card_id, card_name
	FROM %s WHERE id = $1`, expensesView)

	row := e.database.QueryRowContext(ctx, selectStmt, id)
	if row.Err() != nil {
		return models.ExpenseView{}, fmt.Errorf("could not query select expenses view by id statement: %v", row.Err())
	}

	var exp models.ExpenseView
	err := row.Scan(
		&exp.Value,
		&exp.Date,
		&exp.Description,
		&exp.CategoryID,
		&exp.Category,
		&exp.SubCategoryID,
		&exp.SubCategory,
		&exp.CardID,
		&exp.Card,
	)
	if err != nil {
		return models.ExpenseView{}, fmt.Errorf("could not scan expense fields in get expense by id: %v", err)
	}

	exp.ID = id

	return exp, nil
}

// GetExpensesByDates gets expenses from the expenses db table that matches the dates' range provided
func (e ExpensesRepo) GetExpensesByDates(
	ctx context.Context,
	minDate time.Time,
	maxDate time.Time,
) ([]models.ExpenseView, error) {

	selectStmt := fmt.Sprintf(`SELECT 
	value, date, description, category_id, category_name, 
	subcategory_id, subcategory_name, card_id, card_name
	FROM %s WHERE date BETWEEN $1 AND $2`, expensesView)

	rows, err := e.database.QueryContext(ctx, selectStmt, minDate, maxDate)
	if err != nil {
		return []models.ExpenseView{}, fmt.Errorf("could not query select expenses view by dates statement: %v", err)
	}
	defer rows.Close()

	var expenses []models.ExpenseView

	var exp models.ExpenseView

	for rows.Next() {
		err := rows.Scan(
			&exp.Value,
			&exp.Date,
			&exp.Description,
			&exp.CategoryID,
			&exp.Category,
			&exp.SubCategoryID,
			&exp.SubCategory,
			&exp.CardID,
			&exp.Card,
		)
		if err != nil {
			return []models.ExpenseView{}, fmt.Errorf("could not scan expense fields in get expenses by dates: %v", err)
		}

		expenses = append(expenses, exp)
	}

	err = rows.Err()
	if err != nil {
		return []models.ExpenseView{},
			fmt.Errorf("found error after scanning all expenses fields in get expenses by dates: %v", err)
	}

	return expenses, nil
}

// GetExpensesByCategory gets expenses from the expenses db table that matches the category provided
func (e ExpensesRepo) GetExpensesByCategory(ctx context.Context, category string) ([]models.ExpenseView, error) {

	selectStmt := fmt.Sprintf(`SELECT 
	value, date, description, category_id, category_name, 
	subcategory_id, subcategory_name, card_id, card_name
	FROM %s WHERE category_name = $1`, expensesView)

	rows, err := e.database.QueryContext(ctx, selectStmt, category)
	if err != nil {
		return []models.ExpenseView{}, fmt.Errorf("could not query select expenses view by category statement: %v", err)
	}
	defer rows.Close()

	var expenses []models.ExpenseView

	var exp models.ExpenseView

	for rows.Next() {
		err := rows.Scan(
			&exp.Value,
			&exp.Date,
			&exp.Description,
			&exp.CategoryID,
			&exp.Category,
			&exp.SubCategoryID,
			&exp.SubCategory,
			&exp.CardID,
			&exp.Card,
		)
		if err != nil {
			return []models.ExpenseView{}, fmt.Errorf("could not scan expense fields in get expenses by category: %v", err)
		}

		expenses = append(expenses, exp)
	}

	err = rows.Err()
	if err != nil {
		return []models.ExpenseView{},
			fmt.Errorf("found error after scanning all expenses fields in get expenses by category: %v", err)
	}

	return expenses, nil
}

// GetExpensesBySubCategory gets expenses from the expenses db table that matches the subcategory provided
func (e ExpensesRepo) GetExpensesBySubCategory(ctx context.Context, subCategory string) ([]models.ExpenseView, error) {

	selectStmt := fmt.Sprintf(`SELECT 
	value, date, description, category_id, category_name, 
	subcategory_id, subcategory_name, card_id, card_name
	FROM %s WHERE subcategory_name = $1`, expensesView)

	rows, err := e.database.QueryContext(ctx, selectStmt, subCategory)
	if err != nil {
		return []models.ExpenseView{}, fmt.Errorf("could not query select expenses view by subcategory statement: %v", err)
	}
	defer rows.Close()

	var expenses []models.ExpenseView

	var exp models.ExpenseView

	for rows.Next() {
		err := rows.Scan(
			&exp.Value,
			&exp.Date,
			&exp.Description,
			&exp.CategoryID,
			&exp.Category,
			&exp.SubCategoryID,
			&exp.SubCategory,
			&exp.CardID,
			&exp.Card,
		)
		if err != nil {
			return []models.ExpenseView{}, fmt.Errorf("could not scan expense fields in get expenses by subategory: %v", err)
		}

		expenses = append(expenses, exp)
	}

	err = rows.Err()
	if err != nil {
		return []models.ExpenseView{},
			fmt.Errorf("found error after scanning all expenses fields in get expenses by subcategory: %v", err)
	}

	return expenses, nil
}

// GetExpensesByCard gets expenses from the expenses db table that matches the card provided
func (e ExpensesRepo) GetExpensesByCard(ctx context.Context, card string) ([]models.ExpenseView, error) {

	selectStmt := fmt.Sprintf(`SELECT 
	value, date, description, category_id, category_name, 
	subcategory_id, subcategory_name, card_id, card_name
	FROM %s WHERE card_name = $1`, expensesView)

	rows, err := e.database.QueryContext(ctx, selectStmt, card)
	if err != nil {
		return []models.ExpenseView{}, fmt.Errorf("could not query select expenses by card statement: %v", err)
	}
	defer rows.Close()

	var expenses []models.ExpenseView

	var exp models.ExpenseView

	for rows.Next() {
		err := rows.Scan(
			&exp.Value,
			&exp.Date,
			&exp.Description,
			&exp.CategoryID,
			&exp.Category,
			&exp.SubCategoryID,
			&exp.SubCategory,
			&exp.CardID,
			&exp.Card,
		)
		if err != nil {
			return []models.ExpenseView{}, fmt.Errorf("could not scan expense fields in get expenses by card: %v", err)
		}

		expenses = append(expenses, exp)
	}

	err = rows.Err()
	if err != nil {
		return []models.ExpenseView{}, fmt.Errorf("found error after scanning all expenses fields in get expenses by card: %v", err)
	}

	return expenses, nil
}

// DeleteExpense deletes an expense from the expenses db table
func (e ExpensesRepo) DeleteExpense(ctx context.Context, id int64) error {

	deleteStmt := fmt.Sprintf(`DELETE FROM %s 
	WHERE id = $1`, expensesTable)

	result, err := e.database.ExecContext(ctx, deleteStmt, id)
	if err != nil {
		return fmt.Errorf("could not exec expense delete statement: %v", err)
	}

	numRowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get number of rows affected in exec expense delete statement: %v", err)
	}

	if numRowsAffected == 0 {
		return ErrNoRowsAffectedExpenseDelete
	}

	return nil
}
