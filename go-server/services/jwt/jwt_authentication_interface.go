package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

//jwt service
type JWTInterface interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}
