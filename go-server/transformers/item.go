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

func GetItemListResponse(items []entities.Items) models.ItemList {
	itemList := []models.Item{}
	for _, v := range items {
		item := models.Item{
			ID:   v.ID,
			Name: v.Name,
		}
		itemList = append(itemList, item)
	}
	return models.ItemList{
		Item: itemList,
	}
}
