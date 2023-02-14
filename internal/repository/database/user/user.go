package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

const (
	tableNameUsers = "users"
)

// DB implements the user repository methods
type DB struct {
	database *sql.DB
}

// NewDB creates a new UserRepo
func NewDB(database *sql.DB) DB {
	return DB{
		database: database,
	}
}

func (u DB) InsertUser(ctx context.Context, user models.UserTable) (int64, error) {

	insertStmt := fmt.Sprintf("INSERT INTO %s (username, passhash) VALUES ($1, $2) RETURNING id", tableNameUsers)

	var id int64

	err := u.database.QueryRowContext(ctx, insertStmt, user.Username, user.Passhash).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error scanning user id: %v", err)
	}

	return id, nil
}

func (u DB) GetUserByUsername(ctx context.Context, username string) (models.UserTable, error) {

	selectStmt := fmt.Sprintf("SELECT id, username, passhash FROM %s WHERE username = $1", tableNameUsers)

	row := u.database.QueryRowContext(ctx, selectStmt, username)

	var user models.UserTable
	err := row.Scan(&user.ID, &user.Username, &user.Passhash)
	if err != nil {
		return models.UserTable{}, fmt.Errorf("error scanning user fields: %v", err)
	}

	return user, nil
}
