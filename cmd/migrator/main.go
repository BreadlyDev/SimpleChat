package main

import (
	"errors"
	"fmt"
	"log"
	"simplechat/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var migrationsPath, migrationsTable, DBUrl string
	migrationsPath = "./migrations"
	migrationsTable = "migrations"

	cfg := &config.MustLoad().Storage

	DBUrl = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&x-migrations-table=%s",
		cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.DBName, migrationsTable)

	m, err := migrate.New(
		"file://"+migrationsPath,
		DBUrl,
	)

	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return
		}

		panic(err)
	}

	log.Println("migrations applied")
}
