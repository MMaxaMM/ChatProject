package service

import (
	"chat/internal/models"
	"chat/internal/repository"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	rep *repository.Repository
}

func NewAuthService(rep *repository.Repository) *AuthService {
	return &AuthService{rep: rep}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}

func (s *AuthService) CreateUser(request *models.SignUpRequest) (*models.SignUpResponse, error) {
	const op = "service.CreateUser"

	passwordHash := generatePasswordHash(request.Password)
	userID, err := s.rep.CreateUser(request.Username, passwordHash)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	response := models.SignUpResponse{UserId: userID}
	return &response, nil
}

func (s *AuthService) GenerateToken(request *models.SignInRequest) (*models.SignInResponse, error) {
	const op = "service.GenerateToken"

	passwordHash := generatePasswordHash(request.Password)
	userId, err := s.rep.GetUserId(request.Username, passwordHash)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})

	signedToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	response := models.SignInResponse{Token: signedToken}

	return &response, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
