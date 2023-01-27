package mock

import (
	"context"
	"errors"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

var (
	IncomeSalaryCard = models.CardTable{
		ID:   1,
		Name: "CGD",
	}
)

// Card mocks the card repository methods
type Card struct {
}

// NewCard creates a Card mock
func NewCard() Card {
	return Card{}
}

// InsertCard mocks a card insert
func (c Card) InsertCard(ctx context.Context, card models.CardTable) (int64, error) {

	switch card {
	case IncomeSalaryCard:
		return 1, nil
	default:
		return 0, errors.New("could not insert card")
	}
}

// UpdateCard mocks a card update
func (c Card) UpdateCard(ctx context.Context, card models.CardTable) (int64, error) {

	switch card {
	case IncomeSalaryCard:
		return 1, nil
	default:
		return 0, errors.New("could not update card")
	}
}

// GetCardByID returns the card from the cache if card with that id exists
func (c Card) GetCardByID(ctx context.Context, id int64) (models.CardTable, error) {

	switch id {
	case IncomeSalaryCard.ID:
		return IncomeSalaryCard, nil
	default:
		return models.CardTable{}, errors.New("could not get card by id")
	}
}

// GetCardByName returns the card from the cache if card with that name exists
func (c Card) GetCardByName(ctx context.Context, name string) (models.CardTable, error) {

	switch name {
	case IncomeSalaryCard.Name:
		return IncomeSalaryCard, nil
	default:
		return models.CardTable{}, errors.New("could not get card by name")
	}
}

// DeleteCard deletes the card from cache if it exists
func (c Card) DeleteCard(ctx context.Context, id int64) error {

	switch id {
	case IncomeSalaryCard.ID:
		return nil
	default:
		return errors.New("card with this id does not exist")
	}
}
