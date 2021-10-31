package services

import (
	"encoding/json"

	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/providers/repositories"
	"github.com/DeepikaAzad/go-to-do-app/go-server/utils"
	"github.com/gin-gonic/gin"
)

func AddNameInCardItemIds(items string, ctx *gin.Context, arg ...interface{}) ([]models.Item, error) {
	jsonItem := []string{}
	err := json.Unmarshal([]byte(items), &jsonItem)
	resp := []models.Item{}
	if err != nil {
		return resp, err
	}
	resp, err = getItemName(jsonItem, ctx)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func getItemName(itemsObj []string, ctx *gin.Context) ([]models.Item, error) {
	items := []uint{}
	for _, v := range itemsObj {
		items = append(items, utils.ParseUint(v))
	}
	itemsNameList, err := repositories.Items.GetItemsByIds(items, ctx)
	itemsName := []models.Item{}
	if err != nil {
		return itemsName, err
	}
	for _, v := range itemsNameList {
		itemsName = append(itemsName, models.Item{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	return itemsName, nil
}
