package expense

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rubengomes8/golang-personal-finances/internal/models"
)

const (
	tableNameSubCategories = "expense_subcategories"
)

type ExpenseSubCategoryRepo struct {
	database *sql.DB
}

func NewExpenseSubCategoryRepo(database *sql.DB) ExpenseSubCategoryRepo {
	return ExpenseSubCategoryRepo{
		database: database,
	}
}

func (es *ExpenseSubCategoryRepo) InsertExpenseSubCategory(ctx context.Context, expenseSubCategory models.ExpenseSubCategoryTable) (int64, error) {

	insertStmt := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", tableNameSubCategories)

	var id int64
	err := es.database.QueryRowContext(ctx, insertStmt, expenseSubCategory.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (es *ExpenseSubCategoryRepo) UpdateExpenseSubCategory(ctx context.Context, expenseSubCategory models.ExpenseSubCategoryTable) (int64, error) {

	updateStmt := fmt.Sprintf("UPDATE %s SET name = $1, category_id = $2 WHERE id = $3", tableNameSubCategories)

	_, err := es.database.ExecContext(ctx, updateStmt, expenseSubCategory.Name, expenseSubCategory.CategoryId, expenseSubCategory.Id)
	if err != nil {
		return 0, fmt.Errorf("error updating expense subcategory: %v", err)
	}

	return expenseSubCategory.Id, nil
}

func (es *ExpenseSubCategoryRepo) GetExpenseSubCategoryByID(ctx context.Context, id int64) (models.ExpenseSubCategoryTable, error) {

	selectStmt := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tableNameSubCategories)

	row := es.database.QueryRowContext(ctx, selectStmt, id)

	var expenseSubCategory models.ExpenseSubCategoryTable
	err := row.Scan(&expenseSubCategory.Id, &expenseSubCategory.Name, &expenseSubCategory.CategoryId)
	if err != nil {
		return models.ExpenseSubCategoryTable{}, err
	}

	return expenseSubCategory, nil
}

func (es *ExpenseSubCategoryRepo) GetExpenseSubCategoryByName(ctx context.Context, name string) (models.ExpenseSubCategoryTable, error) {

	selectStmt := fmt.Sprintf("SELECT * FROM %s WHERE name = $1", tableNameSubCategories)

	row := es.database.QueryRowContext(ctx, selectStmt, name)

	var expenseSubCategory models.ExpenseSubCategoryTable
	err := row.Scan(&expenseSubCategory.Id, &expenseSubCategory.Name, &expenseSubCategory.CategoryId)
	if err != nil {
		return models.ExpenseSubCategoryTable{}, err
	}

	return expenseSubCategory, nil
}

func (es *ExpenseSubCategoryRepo) DeleteExpenseSubCategory(ctx context.Context, id int64) error {

	deleteStmt := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableNameSubCategories)

	result, err := es.database.ExecContext(ctx, deleteStmt, id)
	if err != nil {
		return fmt.Errorf("error deleting expense subcategory by id: %v", err)
	}

	numRowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get number of rows affected in exec expense subcategory delete statement: %v", err)
	}

	if numRowsAffected == 0 {
		return fmt.Errorf("there were no rows affected in exec expense subcategory delete statement")
	}

	return nil
}
