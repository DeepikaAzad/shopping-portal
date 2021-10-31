package orders

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrdersGorm struct{}

func (r OrdersGorm) CreateOrder(order entities.Orders, ctx *gin.Context) (entities.Orders, error) {
	dbx, _ := ctx.Get("DB")
	db := dbx.(*gorm.DB)
	err := db.Create(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}
