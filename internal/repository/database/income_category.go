package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

const (
	tableNameIncomeCategories = "income_categories"
)

// IncomeCategoryRepo implements the income category repository methods
type IncomeCategoryRepo struct {
	database *sql.DB
}

// NewIncomeCategoryRepo creates a new IncomeCategoryRepo
func NewIncomeCategoryRepo(database *sql.DB) IncomeCategoryRepo {
	return IncomeCategoryRepo{
		database: database,
	}
}

// InsertIncomeCategory inserts an income category on the income categories db table
func (ic IncomeCategoryRepo) InsertIncomeCategory(ctx context.Context, incomeCategory models.IncomeCategoryTable) (int64, error) {

	insertStmt := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", tableNameIncomeCategories)

	var id int64

	err := ic.database.QueryRowContext(ctx, insertStmt, incomeCategory.Name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not scan income category id: %v", err)
	}

	return id, nil
}

// UpdateIncomeCategory updates an income category on the income categories db table
func (ic IncomeCategoryRepo) UpdateIncomeCategory(
	ctx context.Context,
	incomeCategory models.IncomeCategoryTable,
) (int64, error) {

	updateStmt := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", tableNameIncomeCategories)

	_, err := ic.database.ExecContext(ctx, updateStmt, incomeCategory.Name, incomeCategory.ID)
	if err != nil {
		return 0, fmt.Errorf("error updating income category: %v", err)
	}

	return incomeCategory.ID, nil
}

// GetIncomeCategoryByID gets an income category from the income categories db table by id
func (ic IncomeCategoryRepo) GetIncomeCategoryByID(
	ctx context.Context,
	id int64,
) (models.IncomeCategoryTable, error) {

	selectStmt := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tableNameIncomeCategories)

	row := ic.database.QueryRowContext(ctx, selectStmt, id)

	var incomeCategory models.IncomeCategoryTable

	err := row.Scan(&incomeCategory.ID, &incomeCategory.Name)
	if err != nil {
		return models.IncomeCategoryTable{}, fmt.Errorf("could not scan income category fields: %v", err)
	}

	return incomeCategory, nil
}

// GetIncomeCategoryByName gets an income category from the income categories db table by name
func (ic IncomeCategoryRepo) GetIncomeCategoryByName(
	ctx context.Context,
	name string,
) (models.IncomeCategoryTable, error) {

	selectStmt := fmt.Sprintf("SELECT * FROM %s WHERE name = $1", tableNameIncomeCategories)

	row := ic.database.QueryRowContext(ctx, selectStmt, name)

	var incomeCategory models.IncomeCategoryTable
	err := row.Scan(&incomeCategory.ID, &incomeCategory.Name)
	if err != nil {
		return models.IncomeCategoryTable{}, fmt.Errorf("could not scan income category fields: %v", err)
	}

	return incomeCategory, nil
}

// DeleteIncomeCategory deletes an income category from the income categories db table
func (ic IncomeCategoryRepo) DeleteIncomeCategory(
	ctx context.Context,
	id int64,
) error {

	deleteStmt := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableNameIncomeCategories)

	result, err := ic.database.ExecContext(ctx, deleteStmt, id)
	if err != nil {
		return fmt.Errorf("error deleting income category by id: %v", err)
	}

	numRowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get number of rows affected in exec income category delete statement: %v", err)
	}

	if numRowsAffected == 0 {
		return ErrNoRowsAffectedExpCategoryDelete
	}

	return nil
}
