package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

const (
	tableNameUsers = "users"
)

// UserRepo implements the user repository methods
type UserRepo struct {
	database *sql.DB
}

// NewUserRepo creates a new UserRepo
func NewUserRepo(database *sql.DB) UserRepo {
	return UserRepo{
		database: database,
	}
}

func (u UserRepo) InsertUser(ctx context.Context, user models.UserTable) (int64, error) {

	insertStmt := fmt.Sprintf("INSERT INTO %s (username, salt, passhash) VALUES ($1, $2, $3) RETURNING id", tableNameUsers)

	var id int64

	err := u.database.QueryRowContext(ctx, insertStmt, user.Username, user.Salt, user.Passhash).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error scanning user id: %v", err)
	}

	return id, nil
}

func (u UserRepo) GetUserByUsername(ctx context.Context, username string) (models.UserTable, error) {

	selectStmt := fmt.Sprintf("SELECT id, username, salt, passhash FROM %s WHERE name = $1", tableNameUsers)

	row := u.database.QueryRowContext(ctx, selectStmt, username)

	var user models.UserTable
	err := row.Scan(&user.ID, &user.Username, &user.Salt, &user.Passhash)
	if err != nil {
		return models.UserTable{}, fmt.Errorf("error scanning user fields: %v", err)
	}

	return user, nil
}
