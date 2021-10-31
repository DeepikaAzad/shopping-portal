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

func (r OrdersGorm) OrderList(usersId uint, ctx *gin.Context) ([]entities.ItemOrder, error) {
	dbx, _ := ctx.Get("DB")
	db := dbx.(*gorm.DB)
	orders := []entities.Orders{}
	itemOrders := []entities.ItemOrder{}
	rows, err := db.
		Joins("left join carts on carts.id = orders.carts_id").
		Where("orders.users_id = ? and carts.is_purchased = ?", usersId, 1).
		Select("orders.carts_id", "carts.items_id").
		Find(&orders).Rows()

	if err != nil {
		return itemOrders, err
	}

	for rows.Next() {
		order := entities.ItemOrder{}
		if err := rows.Scan(&order.CartID, &order.ItemsID); err != nil {
			return itemOrders, err
		}
		itemOrders = append(itemOrders, order)
	}
	defer rows.Close()
	return itemOrders, nil
}
