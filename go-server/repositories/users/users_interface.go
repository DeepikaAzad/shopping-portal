package users

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
	"github.com/gin-gonic/gin"
)

type UsersGormInterface interface {
	CreateUser(entities.Users, *gin.Context) (entities.Users, error)
	UpdateUser(entities.Users, string, *gin.Context) (entities.Users, error)
	GetUserByUserName(string, *gin.Context) (entities.Users, error)
	GetUserList(*gin.Context) ([]entities.Users, error)
}
