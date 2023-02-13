package tools

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq" //no lint
)

func InitPostgres(localhost string) (*sql.DB, error) {

	dbPortEnv := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(dbPortEnv)
	if err != nil {
		return nil, err
	}

	coonectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		localhost, dbPort, os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", coonectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
