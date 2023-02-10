package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	golangmigrate "github.com/golang-migrate/migrate"
	golangmigratePostgres "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"                //no lint
	_ "github.com/jackc/pgx/stdlib"                                  //no lint
	_ "github.com/rubengomes8/golang-personal-finances/internal/env" //no lint
	"github.com/rubengomes8/golang-personal-finances/internal/tools/postgres"
	"github.com/urfave/cli"
)

func loadSQL(file string) ([]string, error) {
	commentBlock := regexp.MustCompile(`/\*([^*]|[\r\n]|(\*+([^*/]|[\r\n])))*\*+/`)

	data, err := os.ReadFile(file) // nolint
	if err != nil {
		return nil, err
	}

	result := bytes.Split(data, []byte(";"))
	result = result[:len(result)-1] // remove the empty one
	queries := make([]string, len(result))
	for index, value := range result {
		queries[index] = commentBlock.ReplaceAllString(strings.TrimSpace(string(value)+";"), "")
	}

	return queries, nil
}

func runSQL(db *sql.DB, path string) error {
	queries, err := loadSQL(path)
	if err != nil {
		return err
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}

func extensionsDB(*cli.Context) error {
	db, err := postgres.Init()
	if err != nil {
		return err
	}

	return runSQL(db, os.Getenv("DB_EXTENSIONS_FILEPATH"))
}

func instance(pool *sql.DB, source string) (*golangmigrate.Migrate, error) {
	driver, err := golangmigratePostgres.WithInstance(pool, &golangmigratePostgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := golangmigrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", source),
		"postgres", driver,
	)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func buildInstance() (*golangmigrate.Migrate, error) {
	db, err := postgres.Init()
	if err != nil {
		return nil, err
	}

	return instance(db, os.Getenv("DB_MIGRATIONS_PATH"))
}

func migrateDB(*cli.Context) error {
	m, err := buildInstance()
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != golangmigrate.ErrNoChange {
		return err
	}

	return nil
}

func rollbackDB(*cli.Context) error {
	m, err := buildInstance()
	if err != nil {
		return err
	}

	if err := m.Down(); err != nil && err != golangmigrate.ErrNoChange {
		return err
	}

	return nil
}

func main() {
	c := cli.NewApp()
	c.Commands = []cli.Command{
		{
			Name:   "extensions",
			Usage:  "extensions",
			Action: extensionsDB,
		},
		{
			Name:   "migrate",
			Usage:  "migrate",
			Action: migrateDB,
		},
		{
			Name:   "rollback",
			Usage:  "rollback",
			Action: rollbackDB,
		},
	}

	err := c.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
