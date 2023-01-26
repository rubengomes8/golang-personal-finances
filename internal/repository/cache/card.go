package cache

import (
	"context"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

// Card implements the card repository methods
type Card struct {
	repository []models.CardTable
}

// NewCard creates a Card cache
func NewCard(repository []models.CardTable) Card {
	return Card{
		repository: repository,
	}
}

// InsertCard inserts a card on the cache if card does not exist
func (c *Card) InsertCard(ctx context.Context, card models.CardTable) (int64, error) {

	existingCard, err := c.GetCardByID(ctx, card.ID)
	if err == nil {
		return 0, CardAlreadyExistsError{
			id: existingCard.ID,
		}
	}

	c.repository = append(c.repository, card)

	return 1, nil
}

// UpdateCard updates a card on the cache if card exists
func (c *Card) UpdateCard(ctx context.Context, updatedCard models.CardTable) (int64, error) {

	for idx, card := range c.repository {
		if card.ID == updatedCard.ID {
			c.repository[idx] = updatedCard
			return updatedCard.ID, nil
		}
	}

	return 0, CardNotFoundByIDError{
		id: updatedCard.ID,
	}
}

// GetCardByID returns the card from the cache if card with that id exists
func (c *Card) GetCardByID(ctx context.Context, id int64) (models.CardTable, error) {

	for _, card := range c.repository {
		if card.ID == id {
			return card, nil
		}
	}

	return models.CardTable{}, CardNotFoundByIDError{
		id: id,
	}
}

// GetCardByName returns the card from the cache if card with that name exists
func (c *Card) GetCardByName(ctx context.Context, name string) (models.CardTable, error) {
	for _, card := range c.repository {
		if card.Name == name {
			return card, nil
		}
	}

	return models.CardTable{}, CardNotFoundByNameError{
		name: name,
	}
}

// DeleteCard deletes the card from cache if it exists
func (c *Card) DeleteCard(ctx context.Context, id int64) error {

	for idx, card := range c.repository {
		if card.ID == id {
			c.repository = append(c.repository[:idx], c.repository[idx+1:]...)
			return nil
		}
	}

	return CardNotFoundByIDError{
		id: id,
	}
}
