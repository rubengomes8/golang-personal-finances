package repository

import (
	"context"

	"github.com/rubengomes8/golang-personal-finances/internal/models"
)

type CardRepo interface {
	InsertCard(context.Context, models.Card) (int64, error)
	UpdateCard(context.Context, models.Card) (int64, error)
	GetCardByID(context.Context, int64) (models.Card, error)
	GetCardByName(context.Context, string) (models.Card, error)
	DeleteCard(context.Context, int64) error
}
