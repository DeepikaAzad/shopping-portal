package items

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
	"github.com/gin-gonic/gin"
)

type ItemsInterface interface {
	AddItem(models.AddItemReq, *gin.Context) (entities.Items, error)
	ItemList(*gin.Context) ([]entities.Items, error)
}
