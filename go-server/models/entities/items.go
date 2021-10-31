package entities

import (
	"gorm.io/gorm"
)

type Items struct {
	gorm.Model
	Carts []Carts `gorm:"many2many:carts_item;"`
	Name  string  `gorm:"column:name;type:varchar(50)"`
}
