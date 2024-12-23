package services

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/MMaxaMM/ChatProject/Auth/internal/domain/models"
	"github.com/MMaxaMM/ChatProject/Auth/internal/storage"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	log      *slog.Logger
	storage  AuthStorage
	tokenTTL time.Duration
}

type AuthStorage interface {
	SaveUser(
		ctx context.Context,
		username string,
		passHash []byte,
	) (userID int64, err error)
	User(ctx context.Context, username string) (*models.User, error)
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
)

const secret = "rehdft;br"

// New returns a new interface of the Auth service
func New(
	log *slog.Logger,
	storage AuthStorage,
	tokenTTL time.Duration,
) *AuthService {
	return &AuthService{
		log:      log,
		storage:  storage,
		tokenTTL: tokenTTL,
	}
}

// Login checks if user with given credentials exists in the system and returns access token
//
// If user exists, but password incorrect, returns error
// If user doesn't exists, returns error
func (a *AuthService) Login(
	ctx context.Context,
	username string,
	password string,
) (string, error) {
	const op = "services.Login"
	log := a.log.With(slog.String("op", op), slog.String("username", username))

	log.Info("Attempting to login user")

	user, err := a.storage.User(ctx, username)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Warn("User not found", slog.String("err", err.Error()))
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		log.Error("Failed to get user", slog.String("err", err.Error()))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		log.Info("Invalid credentials", slog.String("err", err.Error()))
		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	log.Info("User logged in successfully")

	token, err := NewToken(user, a.tokenTTL)
	if err != nil {
		log.Error("Failed to generate token", slog.String("err", err.Error()))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

func NewToken(user *models.User, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["user_id"] = user.ID

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Register registers new user in the system and returns user ID
//
// If user with given username already exists, returns error
func (a *AuthService) Register(ctx context.Context,
	username string,
	password string,
) (int64, error) {
	const op = "services.Register"
	log := a.log.With(slog.String("op", op), slog.String("username", username))

	log.Info("Registering user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Failed to generate password hash", slog.String("err", err.Error()))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	userID, err := a.storage.SaveUser(ctx, username, passHash)
	if err != nil {
		if errors.Is(err, storage.ErrUserExists) {
			log.Warn("User already exists", slog.String("err", err.Error()))
			return 0, fmt.Errorf("%s: %w", op, ErrUserExists)
		}
		log.Error("Failed to save user", slog.String("err", err.Error()))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("User registered")

	return userID, nil
}
