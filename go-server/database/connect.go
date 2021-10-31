package database

import (
	"database/sql"
	"log"

	"github.com/DeepikaAzad/go-to-do-app/go-server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connection gets connection of postgresql database
func Connection() (db *gorm.DB, sqlDb *sql.DB) {
	dsn := config.Db.Username + ":" +
		config.Db.Password + "@" +
		config.Db.Connection + "(" +
		config.Db.Host + ":" +
		config.Db.Port + ")/" +
		config.Db.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err, "DB_Connection_Error")
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Panic(err, "DB_Connection_Error")
	}
	return db, sqlDB
}
