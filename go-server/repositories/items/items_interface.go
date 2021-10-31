package items

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
	"github.com/gin-gonic/gin"
)

type ItemGormInterface interface {
	CreateItem(entities.Items, *gin.Context) (entities.Items, error)
	GetItemList(*gin.Context) ([]entities.Items, error)
	GetItemByName(string, *gin.Context) (entities.Items, error)
	GetItemsByIds([]uint, *gin.Context) ([]entities.Items, error)
}
