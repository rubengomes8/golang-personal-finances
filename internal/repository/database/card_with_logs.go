package database

import (
	"context"
	"log"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

type CardWithLogs struct {
	repo Card
}

func NewCardWithLogs(repo Card) CardWithLogs {
	return CardWithLogs{
		repo: repo,
	}
}

func (c CardWithLogs) InsertCard(ctx context.Context, card models.CardTable) (int64, error) {
	log.Printf("Inserting card: %#v\n", card)
	return c.repo.InsertCard(ctx, card)
}

func (c CardWithLogs) UpdateCard(ctx context.Context, card models.CardTable) (int64, error) {
	return c.repo.UpdateCard(ctx, card)
}

func (c CardWithLogs) GetCardByID(ctx context.Context, id int64) (models.CardTable, error) {
	return c.repo.GetCardByID(ctx, id)
}

func (c CardWithLogs) GetCardByName(ctx context.Context, card string) (models.CardTable, error) {
	return c.repo.GetCardByName(ctx, card)
}

func (c CardWithLogs) DeleteCard(ctx context.Context, id int64) error {
	return c.repo.DeleteCard(ctx, id)
}
