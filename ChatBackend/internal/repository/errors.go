package repository

import (
	"chat"

	"github.com/lib/pq"
)

func PostgresNewError(err error) error {
	if err, ok := err.(*pq.Error); ok {
		switch err.Code {
		case "23505":
			return chat.ErrUserDuplicate
		case "23503":
			return chat.ErrForeignKey
		}
	}
	return err
}
