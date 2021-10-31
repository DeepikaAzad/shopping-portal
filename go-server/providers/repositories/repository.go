package repositories

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/repositories/items"
	"github.com/DeepikaAzad/go-to-do-app/go-server/repositories/users"
)

var Users users.UsersGormInterface = users.UsersGorm{}
var Items items.ItemGormInterface = items.ItemsGorm{}
