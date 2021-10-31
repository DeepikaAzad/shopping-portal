package entities

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(50)"`
	UserName string `gorm:"column:user_name;not null;unique;type:varchar(100)"`
	Password string `gorm:"column:password;type:varchar(1000)"`
	Token    string `gorm:"token;type:varchar(1000)"`
	CartsID  uint   `gorm:"carts_id"`
	Carts    Carts
}
