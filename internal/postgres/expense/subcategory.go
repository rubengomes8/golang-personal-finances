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
	// TODO
	return 2, nil
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
	// TODO
	return nil
}
