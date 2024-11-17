package repository

import (
	"chat"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(username, password string) (int64, error) {
	const op = "repository.CreateUser"

	var userId int64
	query := fmt.Sprintf("INSERT INTO %s (username, password) values ($1, $2) RETURNING id", usersTable)

	row := r.db.QueryRow(query, username, password)
	if err := row.Scan(&userId); err != nil {
		return 0, fmt.Errorf("%s: %w", op, PostgresNewError(err))
	}

	return userId, nil
}

func (r *AuthPostgres) GetUserId(username, password string) (int64, error) {
	const op = "repository.GetUserId"

	var userId int64

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", usersTable)
	err := r.db.Get(&userId, query, username, password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("%s: %w", op, chat.ErrUserNotFound)
		}
	}

	return userId, nil
}
