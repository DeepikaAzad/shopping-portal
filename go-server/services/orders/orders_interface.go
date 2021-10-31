package orders

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/gin-gonic/gin"
)

type OrdersInterface interface {
	OrderList(models.OrderListReq, *gin.Context) ([]models.Order, error)
}
