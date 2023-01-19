package postgres

import (
	"database/sql"
	"fmt"
)

// NewDB opens a connection with the database.
func NewDB(host, user, password, name string, port int) (*sql.DB, error) {
	coonectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name)

	return sql.Open("postgres", coonectionString)
}
