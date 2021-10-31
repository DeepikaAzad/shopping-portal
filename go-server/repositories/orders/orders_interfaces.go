package orders

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
	"github.com/gin-gonic/gin"
)

type OrdersGormInterface interface {
	CreateOrder(entities.Orders, *gin.Context) (entities.Orders, error)
	OrderList(uint, *gin.Context) ([]entities.ItemOrder, error)
}
