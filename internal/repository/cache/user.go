package cache

import (
	"context"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

// User implements the user repository methods
type User struct {
	repository []models.UserTable
}

// NewUser creates a Card cache
func NewUser() User {
	return User{
		repository: []models.UserTable{},
	}
}

// InsertUser inserts a user on the cache if user does not exist
func (u *User) InsertUser(ctx context.Context, user models.UserTable) (int64, error) {

	existingUser, err := u.GetUserByUsername(ctx, user.Username)
	if err == nil {
		return 0, UserAlreadyExistsError{
			username: existingUser.Username,
		}
	}

	u.repository = append(u.repository, user)

	return 1, nil
}

// GetUserByUsername returns the user from the cache if user with that name exists
func (u *User) GetUserByUsername(ctx context.Context, username string) (models.UserTable, error) {

	for _, user := range u.repository {
		if user.Username == username {
			return user, nil
		}
	}

	return models.UserTable{}, UserNotFoundByNameError{
		username: username,
	}
}
