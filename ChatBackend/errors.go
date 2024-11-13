package chat

import "errors"

// Handler:

// Service:
var (
	ErrServiceNotAvailable = errors.New("service not available")
)

// Repository:
var (
	ErrUserDuplicate = errors.New("user with that username already exists")
	ErrForeignKey    = errors.New("non-existent foreign key")
	ErrUserNotFound  = errors.New("incorrect username or password")
)
