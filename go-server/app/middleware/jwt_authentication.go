package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/DeepikaAzad/go-to-do-app/go-server/constants"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/services/jwt"
	jwtPkg "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, models.SZLError{
				Type: constants.ErrorCode.INVALID_TOKEN,
			})
			c.Abort()
			return
		}
		tokenString := strings.TrimSpace(authHeader[len(BEARER_SCHEMA):])
		token, err := jwt.JWTAuthService().ValidateToken(tokenString, c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.SZLError{
				Type: constants.ErrorCode.INVALID_TOKEN,
			})
			c.Abort()
			return
		}
		if token.Valid {
			claims := token.Claims.(jwtPkg.MapClaims)
			log.Println(claims)
		} else {
			c.JSON(http.StatusUnauthorized, models.SZLError{
				Type: constants.ErrorCode.INVALID_TOKEN,
			})
			c.Abort()
			return
		}
	}
}
