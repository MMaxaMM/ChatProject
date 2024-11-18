package models

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	PassHash []byte `db:"pass_hash"`
}
