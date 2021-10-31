package orders

import (
	"encoding/json"

	"github.com/DeepikaAzad/go-to-do-app/go-server/constants"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/providers/repositories"
	"github.com/DeepikaAzad/go-to-do-app/go-server/utils"
	"github.com/gin-gonic/gin"
)

type OrdersImpl struct{}

func (i OrdersImpl) OrderList(req models.OrderListReq, ctx *gin.Context) ([]models.Order, error) {
	orderList := []models.Order{}
	orders, err := repositories.Orders.OrderList(req.UsersID, ctx)
	if err != nil {
		return orderList, err
	}

	for _, v := range orders {
		orderJson := []string{}
		err := json.Unmarshal([]byte(v.ItemsID), &orderJson)

		if err != nil {
			return orderList, err
		}
		resp, err := getItemName(orderJson, ctx)
		if err != nil {
			return orderList, err
		}
		order := models.Order{
			CartID:    v.CartID,
			CreatedAt: v.CreatedAt.Format(constants.DateTimeLayout.DD_MM_YYYY_HH_MM_SS),
			ItemNames: resp,
		}
		orderList = append(orderList, order)
	}
	return orderList, nil
}

func getItemName(itemsObj []string, ctx *gin.Context) ([]models.ItemName, error) {
	items := []uint{}
	for _, v := range itemsObj {
		items = append(items, utils.ParseUint(v))
	}
	itemsNameList, err := repositories.Items.GetItemsByIds(items, ctx)
	itemsName := []models.ItemName{}
	if err != nil {
		return itemsName, err
	}
	for _, v := range itemsNameList {
		itemsName = append(itemsName, models.ItemName{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	return itemsName, nil
}
