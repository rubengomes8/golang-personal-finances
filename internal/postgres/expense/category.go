package expense

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rubengomes8/golang-personal-finances/internal/enums"
	models "github.com/rubengomes8/golang-personal-finances/internal/models/rds"
)

const (
	tableNameCategories = "expense_categories"
)

// CategoryRepo implements the expense category repository methods
type CategoryRepo struct {
	database *sql.DB
}

// NewCategoryRepo creates a new CategoryRepo
func NewCategoryRepo(database *sql.DB) CategoryRepo {
	return CategoryRepo{
		database: database,
	}
}

// InsertExpenseCategory inserts an expense category on the expense categories rds table
func (ec *CategoryRepo) InsertExpenseCategory(
	ctx context.Context,
	expenseCategory models.ExpenseCategoryTable,
) (int64, error) {

	insertStmt := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", tableNameCategories)

	var id int64

	err := ec.database.QueryRowContext(ctx, insertStmt, expenseCategory.Name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not scan expense category id: %v", err)
	}

	return id, nil
}

// UpdateExpenseCategory updates an expense category on the expense categories rds table
func (ec *CategoryRepo) UpdateExpenseCategory(
	ctx context.Context,
	expenseCategory models.ExpenseCategoryTable,
) (int64, error) {

	updateStmt := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", tableNameCategories)

	_, err := ec.database.ExecContext(ctx, updateStmt, expenseCategory.Name, expenseCategory.ID)
	if err != nil {
		return 0, fmt.Errorf("error updating expense category: %v", err)
	}

	return expenseCategory.ID, nil
}

// GetExpenseCategoryByID gets an expense category from the expense categories rds table by id
func (ec *CategoryRepo) GetExpenseCategoryByID(
	ctx context.Context,
	id int64,
) (models.ExpenseCategoryTable, error) {

	selectStmt := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tableNameCategories)

	row := ec.database.QueryRowContext(ctx, selectStmt, id)

	var expenseCategory models.ExpenseCategoryTable

	err := row.Scan(&expenseCategory.ID, &expenseCategory.Name)
	if err != nil {
		return models.ExpenseCategoryTable{}, fmt.Errorf("could not scan expense category fields: %v", err)
	}

	return expenseCategory, nil
}

// GetExpenseCategoryByName gets an expense category from the expense categories rds table by name
func (ec *CategoryRepo) GetExpenseCategoryByName(
	ctx context.Context,
	name string,
) (models.ExpenseCategoryTable, error) {

	selectStmt := fmt.Sprintf("SELECT * FROM %s WHERE name = $1", tableNameCategories)

	row := ec.database.QueryRowContext(ctx, selectStmt, name)

	var expenseCategory models.ExpenseCategoryTable
	err := row.Scan(&expenseCategory.ID, &expenseCategory.Name)
	if err != nil {
		return models.ExpenseCategoryTable{}, fmt.Errorf("could not scan expense category fields: %v", err)
	}

	return expenseCategory, nil
}

// DeleteExpenseCategory deletes an expense category from the expense categories rds table
func (ec *CategoryRepo) DeleteExpenseCategory(
	ctx context.Context,
	id int64,
) error {

	deleteStmt := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableNameCategories)

	result, err := ec.database.ExecContext(ctx, deleteStmt, id)
	if err != nil {
		return fmt.Errorf("error deleting expense category by id: %v", err)
	}

	numRowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get number of rows affected in exec expense category delete statement: %v", err)
	}

	if numRowsAffected == 0 {
		return enums.ErrNoRowsAffectedExpCategoryDelete
	}

	return nil
}
