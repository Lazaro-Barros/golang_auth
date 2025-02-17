package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// ITokenService define os métodos para gerar e validar tokens JWT
type ITokenService interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (string, error)
}

type TokenService struct {
	secretKey []byte
}

func NewTokenService() *TokenService {
	return &TokenService{
		secretKey: []byte(os.Getenv("JWT_SECRET")),
	}
}

func (t *TokenService) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(8 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(t.secretKey)
}

// ValidateToken verifica se o token JWT é válido e retorna o userID
func (t *TokenService) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return t.secretKey, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("token inválido ou expirado")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("falha ao ler claims do token")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("user_id ausente no token")
	}

	return userID, nil
}
