package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"simplechat/internal/domain/models"
	"simplechat/internal/storage"

	"golang.org/x/crypto/bcrypt"
)

func (s *Storage) CreateUser(ctx context.Context, user models.User) (int64, error) {
	const op = "storage.postgres.CreateUser"

	stmt, err := s.db.Prepare("INSERT INTO users(username, email, pass_hash) VALUES (?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.ExecContext(ctx, user.Username, user.Email, passHash)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) GetUser(ctx context.Context, email string) (models.UserResponse, error) {
	const op = "storage.postgres.GetUser"

	stmt, err := s.db.Prepare("SELECT id, email, pass_hash FROM users WHERE email = ?")
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, email)

	var user models.UserResponse

	err = row.Scan(&user.ID, &user.Email, &user.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.UserResponse{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}

		return models.UserResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
