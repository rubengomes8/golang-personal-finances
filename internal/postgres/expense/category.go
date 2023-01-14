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

func (ec *ExpenseCategoryRepo) InsertExpenseCategory(ctx context.Context, expenseCategory models.ExpenseCategory) (int64, error) {

	insertStmt := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", tableNameCategories)

	var id int64
	err := ec.database.QueryRowContext(ctx, insertStmt, expenseCategory.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ec *ExpenseCategoryRepo) UpdateExpenseCategory(ctx context.Context, expenseCategory models.ExpenseCategory) (int64, error) {
	return 2, nil
}

func (ec *ExpenseCategoryRepo) GetExpenseCategoryByID(ctx context.Context, id int64) (models.ExpenseCategory, error) {

	selectStmt := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tableNameCategories)

	row := ec.database.QueryRowContext(ctx, selectStmt, id)

	var expenseCategory models.ExpenseCategory
	err := row.Scan(&expenseCategory.Id, &expenseCategory.Name)
	if err != nil {
		return models.ExpenseCategory{}, err
	}

	return expenseCategory, nil
}

func (ec *ExpenseCategoryRepo) GetExpenseCategoryByName(ctx context.Context, name string) (models.ExpenseCategory, error) {

	selectStmt := fmt.Sprintf("SELECT * FROM %s WHERE name = $1", tableNameCategories)

	row := ec.database.QueryRowContext(ctx, selectStmt, name)

	var expenseCategory models.ExpenseCategory
	err := row.Scan(&expenseCategory.Id, &expenseCategory.Name)
	if err != nil {
		return models.ExpenseCategory{}, err
	}

	return expenseCategory, nil
}

func (ec *ExpenseCategoryRepo) DeleteExpenseCategory(ctx context.Context, id int64) error {
	return nil
}
