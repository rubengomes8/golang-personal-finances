package service

import (
	"context"

	"github.com/rubengomes8/golang-personal-finances/internal/http/models"
)

type Incomes interface {
	Create(context.Context, models.Income) (int, error)
	Update(context.Context, models.Income) error
	Delete(context.Context, int) error
	GetByID(context.Context, int) (models.Income, error)
	GetAllByCard(context.Context, string) ([]models.Income, error)
	GetAllByCategory(context.Context, string) ([]models.Income, error)
	GetAllByDates(context.Context, string, string) ([]models.Income, error)
}
