package Infrastructure

import (
	"time"

	"task-manager-go/Domain"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secret string
}

func NewJWTService(secret string) Domain.IJWTService {
	return &jwtService{secret: secret}
}

func (j *jwtService) GenerateToken(user Domain.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}
