package tools

import (
	"database/sql"
	"os"
	"strconv"

	_ "github.com/lib/pq" //no lint
	"github.com/rubengomes8/golang-personal-finances/internal/repository/database"
)

func InitPostgres() (*sql.DB, error) {

	dbPortEnv := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(dbPortEnv)
	if err != nil {
		return nil, err
	}

	db, err := database.New(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_NAME"),
		dbPort,
	)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
