package entities

import (
	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	Carts   Carts
	CartsID uint `gorm:"cart_id;foreignkey:carts_id;"`
	Users   Users
	UsersID uint `gorm:"user_id;foreignkey:users_id;"`
}
