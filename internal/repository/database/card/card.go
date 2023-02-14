package card

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

const (
	tableNameCards = "cards"
)

// Database implements the card repository methods
type Database struct {
	database *sql.DB
}

// NewDatabase creates a new CardRepo
func NewDatabase(database *sql.DB) Database {
	return Database{
		database: database,
	}
}

// InsertCard inserts a card on the cards' db table
func (c Database) InsertCard(ctx context.Context, card models.CardTable) (int64, error) {

	insertStmt := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", tableNameCards)

	var id int64

	err := c.database.QueryRowContext(ctx, insertStmt, card.Name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error scanning card id: %v", err)
	}

	return id, nil
}

// UpdateCard updates a card on the cards' db table
func (c Database) UpdateCard(ctx context.Context, card models.CardTable) (int64, error) {
	updateStmt := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", tableNameCards)

	_, err := c.database.ExecContext(ctx, updateStmt, card.Name, card.ID)
	if err != nil {
		return 0, fmt.Errorf("error updating card: %v", err)
	}

	return card.ID, nil
}

// GetCardByID gets a card from the cards' db table by id
func (c Database) GetCardByID(ctx context.Context, id int64) (models.CardTable, error) {

	selectStmt := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tableNameCards)

	row := c.database.QueryRowContext(ctx, selectStmt, id)

	var card models.CardTable

	err := row.Scan(&card.ID, &card.Name)
	if err != nil {
		return models.CardTable{}, fmt.Errorf("error scanning card fields: %v", err)
	}

	return card, nil
}

// GetCardByName gets a card from the cards' db table by name
func (c Database) GetCardByName(ctx context.Context, name string) (models.CardTable, error) {

	selectStmt := fmt.Sprintf("SELECT * FROM %s WHERE name = $1", tableNameCards)

	row := c.database.QueryRowContext(ctx, selectStmt, name)

	var card models.CardTable
	err := row.Scan(&card.ID, &card.Name)
	if err != nil {
		return models.CardTable{}, fmt.Errorf("error scanning card fields: %v", err)
	}

	return card, nil
}

// DeleteCardRepo deletes a card from the cards' db table
func (c Database) DeleteCard(ctx context.Context, id int64) error {
	deleteStmt := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableNameCards)

	result, err := c.database.ExecContext(ctx, deleteStmt, id)
	if err != nil {
		return fmt.Errorf("error deleting card by id: %v", err)
	}

	numRowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get number of rows affected in exec card delete statement: %v", err)
	}

	if numRowsAffected == 0 {
		return ErrNoRowsAffectedOnDelete
	}

	return nil
}
