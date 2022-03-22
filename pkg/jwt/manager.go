package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Manager struct {
	secret  string
	expires time.Duration
}

type UserClaims struct {
	ID uint32 `json:"id"`
	jwt.StandardClaims
}

func (manager *Manager) Generate(id uint32) (string, error) {
	claims := UserClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.expires).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secret))
}

func (manager *Manager) Validate(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, status.Error(codes.Unauthenticated, "invalid token")
			}

			return []byte(manager.secret), nil
		},
	)

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	return claims, nil
}
