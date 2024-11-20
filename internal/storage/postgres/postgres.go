package postgres

import (
	"database/sql"
	"fmt"
	"simplechat/internal/config"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(DBConfig *config.DB) (*Storage, error) {
	const op = "storage.postgres.New"

	strConn := GetDBUrl(DBConfig)

	db, err := sql.Open("postgres", strConn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func GetDBUrl(DBConfig *config.DB) string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		DBConfig.User, DBConfig.Pass, DBConfig.DBName, DBConfig.Host, DBConfig.Port, DBConfig.SSLMode,
	)
}

// Ping is temporal function to check the connetion to DB
func (s *Storage) Ping() error {
	const op = "storage.postgres.Ping"

	err := s.db.Ping()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
