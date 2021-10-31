package providers

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/services/carts"
	"github.com/DeepikaAzad/go-to-do-app/go-server/services/items"
	"github.com/DeepikaAzad/go-to-do-app/go-server/services/orders"
	"github.com/DeepikaAzad/go-to-do-app/go-server/services/users"
)

var Users users.UsersInterface = users.UsersImpl{}
var Items items.ItemsInterface = items.ItemsImpl{}
var Carts carts.CartsInterface = carts.CartsImpl{}
var Orders orders.OrdersInterface = orders.OrdersImpl{}
