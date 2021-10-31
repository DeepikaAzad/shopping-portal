package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//jwt service
type JWTInterface interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string, ctx *gin.Context) (*jwt.Token, error)
}
