package service

import (
	"chat"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type MiddlewareService struct{}

func NewMiddlewareService() *MiddlewareService {
	return &MiddlewareService{}
}

func (s *MiddlewareService) ParseToken(accessToken string) (int64, error) {
	const op = "service.ParseToken"

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		err = errors.New("token claims are not of type *tokenClaims")
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return 0, fmt.Errorf("%s: %w", op, chat.ErrTokenExpired)
	}

	return claims.UserId, nil
}
