package repository

import (
	"context"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

//go:generate gowrap gen -g -i CardRepo -t ./templates/log_template.go.tmpl -o ./database/card/with_logs_by_template.go
//go:generate gowrap gen -g -i CardRepo -t ./templates/red_template.go.tmpl -o ./database/card/with_red_by_template.go
// CardRepo defines the card repository interface.
type CardRepo interface {
	InsertCard(context.Context, models.CardTable) (int64, error)
	UpdateCard(context.Context, models.CardTable) (int64, error)
	GetCardByID(context.Context, int64) (models.CardTable, error)
	GetCardByName(context.Context, string) (models.CardTable, error)
	DeleteCard(context.Context, int64) error
}
