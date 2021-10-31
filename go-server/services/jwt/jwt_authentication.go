package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/DeepikaAzad/go-to-do-app/go-server/providers/repositories"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type JWTImpl struct {
	secretKey string
	issure    string
}

//auth-jwt
func JWTAuthService() JWTInterface {
	return &JWTImpl{
		secretKey: getSecretKey(),
		issure:    "Deepika",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	fmt.Println("SECRET = " + secret)
	return secret
}

func (service *JWTImpl) GenerateToken(name string, isUser bool) string {
	claims := &authCustomClaims{
		name,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *JWTImpl) ValidateToken(encodedToken string, ctx *gin.Context) (*jwt.Token, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(encodedToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(getSecretKey()), nil
	})
	if err != nil {
		return nil, err
	}

	for key, val := range claims {
		if key == "name" {
			fmt.Println(fmt.Sprint(val))
			user, err := repositories.Users.GetUserByUserName(fmt.Sprint(val), ctx)
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}

			if user.Token != encodedToken {
				return nil, errors.New("invalid token")
			}
			continue
		}
	}

	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}
