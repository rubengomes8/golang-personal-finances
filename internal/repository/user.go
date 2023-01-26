package repository

import (
	"context"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

// UserRepo defines the user repository interface.
type UserRepo interface {
	InsertUser(context.Context, models.UserTable) (int64, error)
	GetUserByUsername(context.Context, string) (models.UserTable, error)
}
