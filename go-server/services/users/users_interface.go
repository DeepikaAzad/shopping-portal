package users

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
	"github.com/gin-gonic/gin"
)

type UsersInterface interface {
	RegisterUser(models.RegisterUserReq, *gin.Context) (entities.Users, error)
	LoginUser(models.LoginUserReq, *gin.Context) (entities.Users, error)
}
