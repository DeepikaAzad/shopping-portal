package entities

import "time"

type Items struct {
	ID        uint   `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string `gorm:"column:name;type:varchar(50)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
