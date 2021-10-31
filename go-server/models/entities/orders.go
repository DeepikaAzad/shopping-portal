package entities

import "time"

type Orders struct {
	ID        uint `gorm:"column:id;primaryKey;autoIncrement"`
	Carts     Carts
	CartsID   uint `gorm:"column:carts_id;foreignkey:cart_id;"`
	Users     Users
	UsersID   uint `gorm:"column:users_id;foreignkey:user_id;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
