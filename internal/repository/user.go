package repository

import (
	"context"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

//go:generate gowrap gen -g -i UserRepo -t ./templates/log_template.go.tmpl -o ./database/user/with_logs_by_template.go
//go:generate gowrap gen -g -i UserRepo -t ./templates/red_template.go.tmpl -o ./database/user/with_red_by_template.go
// UserRepo defines the user repository interface.
type UserRepo interface {
	InsertUser(context.Context, models.UserTable) (int64, error)
	GetUserByUsername(context.Context, string) (models.UserTable, error)
}
