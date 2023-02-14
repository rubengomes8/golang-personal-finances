package card

import (
	"context"
	"log"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

type DBWithLogs struct {
	repo Database
}

func NewDBWithLogs(repo Database) DBWithLogs {
	return DBWithLogs{
		repo: repo,
	}
}

func (c DBWithLogs) InsertCard(ctx context.Context, card models.CardTable) (int64, error) {
	log.Printf("Inserting card: %#v\n", card)
	return c.repo.InsertCard(ctx, card)
}

func (c DBWithLogs) UpdateCard(ctx context.Context, card models.CardTable) (int64, error) {
	return c.repo.UpdateCard(ctx, card)
}

func (c DBWithLogs) GetCardByID(ctx context.Context, id int64) (models.CardTable, error) {
	return c.repo.GetCardByID(ctx, id)
}

func (c DBWithLogs) GetCardByName(ctx context.Context, card string) (models.CardTable, error) {
	return c.repo.GetCardByName(ctx, card)
}

func (c DBWithLogs) DeleteCard(ctx context.Context, id int64) error {
	return c.repo.DeleteCard(ctx, id)
}
