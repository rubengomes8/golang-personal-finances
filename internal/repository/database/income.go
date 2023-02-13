package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

const (
	incomesTable = "incomes"
	incomesView  = "incomes_view"
)

// Incomes implements the income repository methods
type Incomes struct {
	database     *sql.DB
	cardRepo     repository.CardRepo
	categoryRepo repository.IncomeCategoryRepo
}

// NewIncomes creates a new Incomes
func NewIncomes(
	database *sql.DB,
	cardRepo repository.CardRepo,
	categoryRepo repository.IncomeCategoryRepo,
) Incomes {
	return Incomes{
		database:     database,
		cardRepo:     cardRepo,
		categoryRepo: categoryRepo,
	}
}

// InsertIncome inserts an income on the incomes db table
func (e Incomes) InsertIncome(ctx context.Context, inc models.IncomeTable) (int64, error) {

	insertStmt := fmt.Sprintf(`INSERT INTO %s 
	(value, date, description, category_id, card_id)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`, incomesTable)

	var id int64

	err := e.database.QueryRowContext(
		ctx,
		insertStmt,
		inc.Value,
		inc.Date,
		inc.Description,
		inc.CategoryID,
		inc.CardID,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not exec income insert statement: %v", err)
	}

	return id, nil
}

// UpdateIncome updates an income on the incomes db table
func (e Incomes) UpdateIncome(ctx context.Context, inc models.IncomeTable) (int64, error) {

	updateStmt := fmt.Sprintf(`UPDATE %s SET 
	(value, date, description, category_id, card_id) =
	($1, $2, $3, $4, $5) WHERE id = $6`, incomesTable)

	result, err := e.database.ExecContext(ctx,
		updateStmt,
		inc.Value,
		inc.Date,
		inc.Description,
		inc.CategoryID,
		inc.CardID,
		inc.ID,
	)
	if err != nil {
		return 0, fmt.Errorf("could not exec income update statement: %v", err)
	}

	numRowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("could not get number of rows affected in exec income update statement: %v", err)
	}

	if numRowsAffected == 0 {
		return 0, ErrNoRowsAffectedIncomeUpdate
	}

	return inc.ID, nil
}

// GetIncomeByID gets an income from the incomes db table by id
func (e Incomes) GetIncomeByID(ctx context.Context, id int64) (models.IncomeView, error) {

	selectStmt := fmt.Sprintf(`SELECT 
	value, date, description, category_id, 
	category_name, card_id, card_name
	FROM %s WHERE id = $1`, incomesView)

	row := e.database.QueryRowContext(ctx, selectStmt, id)
	if row.Err() != nil {
		return models.IncomeView{}, fmt.Errorf("could not query select incomes view by id statement: %v", row.Err())
	}

	var inc models.IncomeView
	err := row.Scan(
		&inc.Value,
		&inc.Date,
		&inc.Description,
		&inc.CategoryID,
		&inc.Category,
		&inc.CardID,
		&inc.Card,
	)
	if err != nil {
		return models.IncomeView{}, fmt.Errorf("could not scan income fields in get income by id: %v", err)
	}

	inc.ID = id

	return inc, nil
}

// GetIncomesByDates gets incomes from the incomes db table that matches the dates' range provided
func (e Incomes) GetIncomesByDates(
	ctx context.Context,
	minDate time.Time,
	maxDate time.Time,
) ([]models.IncomeView, error) {

	selectStmt := fmt.Sprintf(`SELECT 
	value, date, description, category_id,
	category_name, card_id, card_name
	FROM %s WHERE date BETWEEN $1 AND $2`, incomesView)

	rows, err := e.database.QueryContext(ctx, selectStmt, minDate, maxDate)
	if err != nil {
		return []models.IncomeView{}, fmt.Errorf("could not query select incomes view by dates statement: %v", err)
	}
	defer rows.Close()

	var incomes []models.IncomeView

	var inc models.IncomeView

	for rows.Next() {
		err := rows.Scan(
			&inc.Value,
			&inc.Date,
			&inc.Description,
			&inc.CategoryID,
			&inc.Category,
			&inc.CardID,
			&inc.Card,
		)
		if err != nil {
			return []models.IncomeView{}, fmt.Errorf("could not scan income fields in get incomes by dates: %v", err)
		}

		incomes = append(incomes, inc)
	}

	err = rows.Err()
	if err != nil {
		return []models.IncomeView{},
			fmt.Errorf("found error after scanning all incomes fields in get incomes by dates: %v", err)
	}

	return incomes, nil
}

// GetIncomesByCategory gets incomes from the incomes db table that matches the category provided
func (e Incomes) GetIncomesByCategory(ctx context.Context, category string) ([]models.IncomeView, error) {

	selectStmt := fmt.Sprintf(`SELECT 
	value, date, description, category_id, 
	category_name, card_id, card_name
	FROM %s WHERE category_name = $1`, incomesView)

	rows, err := e.database.QueryContext(ctx, selectStmt, category)
	if err != nil {
		return []models.IncomeView{}, fmt.Errorf("could not query select incomes view by category statement: %v", err)
	}
	defer rows.Close()

	var incomes []models.IncomeView

	var inc models.IncomeView

	for rows.Next() {
		err := rows.Scan(
			&inc.Value,
			&inc.Date,
			&inc.Description,
			&inc.CategoryID,
			&inc.Category,
			&inc.CardID,
			&inc.Card,
		)
		if err != nil {
			return []models.IncomeView{}, fmt.Errorf("could not scan income fields in get incomes by category: %v", err)
		}

		incomes = append(incomes, inc)
	}

	err = rows.Err()
	if err != nil {
		return []models.IncomeView{},
			fmt.Errorf("found error after scanning all incomes fields in get incomes by category: %v", err)
	}

	return incomes, nil
}

// GetIncomesByCard gets incomes from the incomes db table that matches the card provided
func (e Incomes) GetIncomesByCard(ctx context.Context, card string) ([]models.IncomeView, error) {

	selectStmt := fmt.Sprintf(`SELECT 
	value, date, description, category_id, 
	category_name, card_id, card_name
	FROM %s WHERE card_name = $1`, incomesView)

	rows, err := e.database.QueryContext(ctx, selectStmt, card)
	if err != nil {
		return []models.IncomeView{}, fmt.Errorf("could not query select incomes by card statement: %v", err)
	}
	defer rows.Close()

	var incomes []models.IncomeView

	var inc models.IncomeView

	for rows.Next() {
		err := rows.Scan(
			&inc.Value,
			&inc.Date,
			&inc.Description,
			&inc.CategoryID,
			&inc.Category,
			&inc.CardID,
			&inc.Card,
		)
		if err != nil {
			return []models.IncomeView{}, fmt.Errorf("could not scan income fields in get incomes by card: %v", err)
		}

		incomes = append(incomes, inc)
	}

	err = rows.Err()
	if err != nil {
		return []models.IncomeView{}, fmt.Errorf("found error after scanning all incomes fields in get incomes by card: %v", err)
	}

	return incomes, nil
}

// DeleteIncome deletes an income from the incomes db table
func (e Incomes) DeleteIncome(ctx context.Context, id int64) error {

	deleteStmt := fmt.Sprintf(`DELETE FROM %s 
	WHERE id = $1`, incomesTable)

	result, err := e.database.ExecContext(ctx, deleteStmt, id)
	if err != nil {
		return fmt.Errorf("could not exec income delete statement: %v", err)
	}

	numRowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get number of rows affected in exec income delete statement: %v", err)
	}

	if numRowsAffected == 0 {
		return ErrNoRowsAffectedIncomeDelete
	}

	return nil
}
