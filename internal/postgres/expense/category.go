package expense

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rubengomes8/golang-personal-finances/internal/models"
)

const (
	tableNameCategories = "expense_categories"
)

type ExpenseCategoryRepo struct {
	database *sql.DB
}

func NewExpenseCategoryRepo(database *sql.DB) ExpenseCategoryRepo {
	return ExpenseCategoryRepo{
		database: database,
	}
}

func (ec *ExpenseCategoryRepo) InsertExpenseCategory(ctx context.Context, expenseCategory models.ExpenseCategoryTable) (int64, error) {

	insertStmt := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", tableNameCategories)

	var id int64
	err := ec.database.QueryRowContext(ctx, insertStmt, expenseCategory.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ec *ExpenseCategoryRepo) UpdateExpenseCategory(ctx context.Context, expenseCategory models.ExpenseCategoryTable) (int64, error) {

	updateStmt := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", tableNameCategories)

	_, err := ec.database.ExecContext(ctx, updateStmt, expenseCategory.Name, expenseCategory.Id)
	if err != nil {
		return 0, fmt.Errorf("error updating expense category: %v", err)
	}

	return expenseCategory.Id, nil
}

func (ec *ExpenseCategoryRepo) GetExpenseCategoryByID(ctx context.Context, id int64) (models.ExpenseCategoryTable, error) {

	selectStmt := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tableNameCategories)

	row := ec.database.QueryRowContext(ctx, selectStmt, id)

	var expenseCategory models.ExpenseCategoryTable
	err := row.Scan(&expenseCategory.Id, &expenseCategory.Name)
	if err != nil {
		return models.ExpenseCategoryTable{}, err
	}

	return expenseCategory, nil
}

func (ec *ExpenseCategoryRepo) GetExpenseCategoryByName(ctx context.Context, name string) (models.ExpenseCategoryTable, error) {

	selectStmt := fmt.Sprintf("SELECT * FROM %s WHERE name = $1", tableNameCategories)

	row := ec.database.QueryRowContext(ctx, selectStmt, name)

	var expenseCategory models.ExpenseCategoryTable
	err := row.Scan(&expenseCategory.Id, &expenseCategory.Name)
	if err != nil {
		return models.ExpenseCategoryTable{}, err
	}

	return expenseCategory, nil
}

func (ec *ExpenseCategoryRepo) DeleteExpenseCategory(ctx context.Context, id int64) error {

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
		return fmt.Errorf("there were no rows affected in exec expense category delete statement")
	}

	return nil
}
