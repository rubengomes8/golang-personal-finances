package expense

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

const (
	tableNameExpenseSubCategories = "expense_subcategories"
)

// SubCategoryDB implements the expense subcategory repository methods
type SubCategoryDB struct {
	database *sql.DB
}

// NewSubCategoryDB creates a new ExpenseSubCategory
func NewSubCategoryDB(database *sql.DB) SubCategoryDB {
	return SubCategoryDB{
		database: database,
	}
}

// InsertExpenseSubCategory inserts an expense subcategory on the expense subcategories db table
func (es SubCategoryDB) InsertExpenseSubCategory(
	ctx context.Context,
	expenseSubCategory models.ExpenseSubCategoryTable,
) (int64, error) {

	insertStmt := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", tableNameExpenseSubCategories)

	var id int64
	err := es.database.QueryRowContext(ctx, insertStmt, expenseSubCategory.Name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not scan expense subcategory id :%v", err)
	}

	return id, nil
}

// UpdateExpenseSubCategory updates an expense subcategory on the expense subcategories db table
func (es SubCategoryDB) UpdateExpenseSubCategory(
	ctx context.Context,
	expenseSubCategory models.ExpenseSubCategoryTable,
) (int64, error) {

	updateStmt := fmt.Sprintf("UPDATE %s SET name = $1, category_id = $2 WHERE id = $3", tableNameExpenseSubCategories)

	_, err := es.database.ExecContext(ctx, updateStmt, expenseSubCategory.Name, expenseSubCategory.CategoryID, expenseSubCategory.ID)
	if err != nil {
		return 0, fmt.Errorf("error updating expense subcategory: %v", err)
	}

	return expenseSubCategory.ID, nil
}

// GetExpenseSubCategoryByID gets an expense subcategory from the expense categories db table by id
func (es SubCategoryDB) GetExpenseSubCategoryByID(
	ctx context.Context,
	id int64,
) (models.ExpenseSubCategoryTable, error) {

	selectStmt := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tableNameExpenseSubCategories)

	row := es.database.QueryRowContext(ctx, selectStmt, id)

	var expenseSubCategory models.ExpenseSubCategoryTable
	err := row.Scan(&expenseSubCategory.ID, &expenseSubCategory.Name, &expenseSubCategory.CategoryID)
	if err != nil {
		return models.ExpenseSubCategoryTable{}, fmt.Errorf("could not scan expense subcategory fields :%v", err)
	}

	return expenseSubCategory, nil
}

// GetExpenseSubCategoryByName gets an expense subcategory from the expense categories db table by name
func (es SubCategoryDB) GetExpenseSubCategoryByName(
	ctx context.Context,
	name string,
) (models.ExpenseSubCategoryTable, error) {

	selectStmt := fmt.Sprintf("SELECT * FROM %s WHERE name = $1", tableNameExpenseSubCategories)

	row := es.database.QueryRowContext(ctx, selectStmt, name)

	var expenseSubCategory models.ExpenseSubCategoryTable
	err := row.Scan(&expenseSubCategory.ID, &expenseSubCategory.Name, &expenseSubCategory.CategoryID)
	if err != nil {
		return models.ExpenseSubCategoryTable{}, fmt.Errorf("could not scan expense subcategory fields :%v", err)
	}

	return expenseSubCategory, nil
}

// DeleteExpenseSubCategory deletes an expense category from the expense subcategories db table
func (es SubCategoryDB) DeleteExpenseSubCategory(ctx context.Context, id int64) error {

	deleteStmt := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableNameExpenseSubCategories)

	result, err := es.database.ExecContext(ctx, deleteStmt, id)
	if err != nil {
		return fmt.Errorf("error deleting expense subcategory by id: %v", err)
	}

	numRowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get number of rows affected in exec expense subcategory delete statement: %v", err)
	}

	if numRowsAffected == 0 {
		return ErrNoRowsAffectedOnSubcategoryDelete
	}

	return nil
}
