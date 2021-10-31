package entities

import "time"

type Carts struct {
	ID          uint   `gorm:"column:id;primaryKey;autoIncrement"`
	UsersID     uint   `gorm:"column:users_id;foreignkey:users_id"`
	ItemsID     string `gorm:"column:items_id"`
	IsPurchased int8   `gorm:"column:is_purchased;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
