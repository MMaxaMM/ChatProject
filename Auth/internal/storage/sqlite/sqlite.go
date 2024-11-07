package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/MMaxaMM/ChatProject/Auth/internal/domain/models"
	"github.com/MMaxaMM/ChatProject/Auth/internal/storage"
	"github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveUser(
	ctx context.Context,
	username string,
	passHash []byte,
) (int64, error) {
	const op = "sqlite.SaveUser"

	stmt, err := s.db.Prepare("INSERT INTO users (username, pass_hash) VALUES (?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.ExecContext(ctx, username, passHash)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrUserExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) User(ctx context.Context, username string) (*models.User, error) {
	op := "sqlite.User"

	stmt, err := s.db.Prepare("SELECT id, pass_hash FROM users WHERE username = ?")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	user := models.User{Username: username}
	row := stmt.QueryRowContext(ctx, username)
	err = row.Scan(&user.ID, &user.PassHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}
