package database

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
	"gorm.io/gorm"
)

type Migrate struct {
	TableName string
	Run       func(*gorm.DB) error
}

func AutoMigrate(db *gorm.DB) []Migrate {
	var users entities.Users
	var items entities.Items
	var carts entities.Carts
	var orders entities.Orders

	cartsM := Migrate{TableName: "carts", Run: func(db *gorm.DB) error {
		db.Migrator().DropTable("carts")
		return db.AutoMigrate(&carts)
	}}
	usersM := Migrate{TableName: "users", Run: func(db *gorm.DB) error {
		db.Migrator().DropTable("users")
		return db.AutoMigrate(&users)
	}}
	itemsM := Migrate{TableName: "items", Run: func(db *gorm.DB) error {
		db.Migrator().DropTable("items")
		return db.AutoMigrate(&items)
	}}
	ordersM := Migrate{TableName: "orders", Run: func(db *gorm.DB) error {
		db.Migrator().DropTable("orders")
		return db.AutoMigrate(&orders)
	}}

	return []Migrate{
		usersM,
		itemsM,
		cartsM,
		ordersM,
	}
}
