package items

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
	"github.com/DeepikaAzad/go-to-do-app/go-server/providers/repositories"
	"github.com/gin-gonic/gin"
)

type ItemsImpl struct{}

func (i ItemsImpl) AddItem(req models.AddItemReq, ctx *gin.Context) (entities.Items, error) {
	itemObj := entities.Items{
		Name: req.Name,
	}
	resp, err := repositories.Items.CreateItem(itemObj, ctx)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
