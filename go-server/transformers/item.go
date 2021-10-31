package transformers

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
)

func GetAddItemResponse(item entities.Items) models.AddItemResp {
	return models.AddItemResp{
		Message: item.Name + " item added sucessfully",
	}
}
