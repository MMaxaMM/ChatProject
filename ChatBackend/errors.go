package chat

import "errors"

// Service error
var (
	ErrServiceNotAvailable = errors.New("service not available")
	ErrTokenExpired        = errors.New("token expired")
)

// Repository error
var (
	ErrUserDuplicate = errors.New("user with that username already exists")
	ErrForeignKey    = errors.New("non-existent foreign key")
	ErrUserNotFound  = errors.New("incorrect username or password")
)
