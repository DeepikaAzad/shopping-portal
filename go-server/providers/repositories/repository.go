package repositories

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/repositories/carts"
	"github.com/DeepikaAzad/go-to-do-app/go-server/repositories/items"
	"github.com/DeepikaAzad/go-to-do-app/go-server/repositories/orders"
	"github.com/DeepikaAzad/go-to-do-app/go-server/repositories/users"
)

var Users users.UsersGormInterface = users.UsersGorm{}
var Items items.ItemGormInterface = items.ItemsGorm{}
var Carts carts.CartGormInterface = carts.CartsGorm{}
var Orders orders.OrdersGormInterface = orders.OrdersGorm{}
