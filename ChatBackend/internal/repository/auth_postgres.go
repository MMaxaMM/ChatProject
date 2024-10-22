package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(username, password string) (int, error) {
	// const op = "repository.chat_interface_postgres.CreateUser"

	var userId int
	query := fmt.Sprintf("INSERT INTO %s (username, password) values ($1, $2) RETURNING id", usersTable)

	row := r.db.QueryRow(query, username, password)
	if err := row.Scan(&userId); err != nil {
		return 0, PostgresNewError(err)
	}

	return userId, nil
}

func (r *AuthPostgres) GetUserId(username, password string) (int, error) {
	// const op = "repository.chat_interface_postgres.GetUserId"

	var userId int
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", usersTable)
	err := r.db.Get(&userId, query, username, password)

	return userId, PostgresNewError(err)
}
