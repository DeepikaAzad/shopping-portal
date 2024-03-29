package entities

import "time"

type Users struct {
	ID        uint   `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string `gorm:"column:name;type:varchar(50)"`
	UserName  string `gorm:"column:user_name;not null;unique;type:varchar(100)"`
	Password  string `gorm:"column:password;not null;type:varchar(1000)"`
	Token     string `gorm:"token;type:varchar(1000)"`
	CartsID   uint   `gorm:"carts_id"`
	Carts     Carts
	CreatedAt time.Time
	UpdatedAt time.Time
}
