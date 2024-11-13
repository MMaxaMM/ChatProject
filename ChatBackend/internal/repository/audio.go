package repository

import "github.com/jmoiron/sqlx"

type AudioPostgres struct {
	db *sqlx.DB
}

func NewAudioPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
