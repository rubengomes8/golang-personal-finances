package card

import (
	"context"
	"database/sql"
	"fmt"

	models "github.com/rubengomes8/golang-personal-finances/internal/models/rds"
)

const (
	tableName = "cards"
)

type CardRepo struct {
	database *sql.DB
}

func NewCardRepo(database *sql.DB) CardRepo {
	return CardRepo{
		database: database,
	}
}

func (c *CardRepo) InsertCard(ctx context.Context, card models.CardTable) (int64, error) {

	insertStmt := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", tableName)

	var id int64
	err := c.database.QueryRowContext(ctx, insertStmt, card.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (c *CardRepo) UpdateCard(ctx context.Context, card models.CardTable) (int64, error) {
	return 2, nil
}

func (c *CardRepo) GetCardByID(ctx context.Context, id int64) (models.CardTable, error) {

	selectStmt := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tableName)

	row := c.database.QueryRowContext(ctx, selectStmt, id)

	var card models.CardTable
	err := row.Scan(&card.Id, &card.Name)
	if err != nil {
		return models.CardTable{}, err
	}

	return card, nil
}

func (c *CardRepo) GetCardByName(ctx context.Context, name string) (models.CardTable, error) {

	selectStmt := fmt.Sprintf("SELECT * FROM %s WHERE name = $1", tableName)

	row := c.database.QueryRowContext(ctx, selectStmt, name)

	var card models.CardTable
	err := row.Scan(&card.Id, &card.Name)
	if err != nil {
		return models.CardTable{}, err
	}

	return card, nil
}

func (c *CardRepo) DeleteCard(ctx context.Context, id int64) error {
	return nil
}
