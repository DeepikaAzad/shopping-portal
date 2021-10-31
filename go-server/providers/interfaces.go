package providers

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/services/items"
	"github.com/DeepikaAzad/go-to-do-app/go-server/services/users"
)

var Users users.UsersInterface = users.UsersImpl{}
var Items items.ItemsInterface = items.ItemsImpl{}
