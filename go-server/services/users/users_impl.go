package users

import (
	"errors"

	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
	"github.com/DeepikaAzad/go-to-do-app/go-server/providers/repositories"
	"github.com/DeepikaAzad/go-to-do-app/go-server/services/jwt"
	"github.com/DeepikaAzad/go-to-do-app/go-server/utils"
	"github.com/gin-gonic/gin"
)

type UsersImpl struct{}

func (u UsersImpl) RegisterUser(user models.RegisterUserReq, ctx *gin.Context) (entities.Users, error) {
	user.Password = utils.DecodeBase64(user.Password)
	userObj := entities.Users{
		Name:     user.FullName,
		UserName: user.UserName,
		Password: user.Password,
	}
	// Hash password
	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return userObj, err
	}
	userObj.Password = password
	resp, err := repositories.Users.CreateUser(userObj, ctx)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (u UsersImpl) LoginUser(user models.LoginUserReq, ctx *gin.Context) (entities.Users, error) {
	user.Password = utils.DecodeBase64(user.Password)
	userObj := entities.Users{
		UserName: user.UserName,
		Password: user.Password, // @TODO:: encrypt
	}
	// Match cred
	registeredUser, err := repositories.Users.GetUserByUserName(userObj.UserName, ctx)
	if err != nil {
		return userObj, err
	}
	token := ""

	// match Password
	if utils.CheckPasswordHash(user.Password, registeredUser.Password) {
		token = jwt.JWTAuthService().GenerateToken(user.UserName, true)
	} else {
		return userObj, errors.New("Invalid password.")
	}

	// update token
	userObj.Token = token
	updatedUser, err := repositories.Users.UpdateUser(userObj, user.UserName, ctx)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}
