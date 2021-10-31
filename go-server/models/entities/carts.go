package entities

import (
	"gorm.io/gorm"
)

type Carts struct {
	gorm.Model
	Name        string `gorm:"column:name;type:varchar(50)"`
	UsersID     uint   `gorm:"column:users_id;foreignkey:users_id"`
	IsPurchased bool   `gorm:"column:is_purchased"`
}
