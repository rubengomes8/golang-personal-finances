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

	date := time.Unix(exp.Date, 0)

	var id int64
	err = e.database.QueryRowContext(ctx, insertStmt, exp.Value, date, exp.Description, category.Id, subCategory.Id, card.Id).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not exec insert expense statement: %v", err)
	}

	return id, nil
}

func (e *ExpenseRepo) UpdateExpense(ctx context.Context, expense models.Expense) (int64, error) {
	return 2, nil
}

func (e *ExpenseRepo) GetExpenseByID(ctx context.Context, id int64) (models.Expense, error) {
	return models.Expense{}, nil
}

func (e *ExpenseRepo) DeleteExpense(ctx context.Context, id int64) error {
	return nil
}
