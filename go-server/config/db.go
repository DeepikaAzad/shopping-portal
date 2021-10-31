package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

//DbConfig ..
type DbConfig struct {
	Connection string `split_words:"true" json:"DB_CONNECTION"`
	Host       string `split_words:"true" json:"DB_HOST"`
	Port       string `split_words:"true" json:"DB_PORT"`
	Database   string `split_words:"true" json:"DB_DATABASE"`
	Username   string `split_words:"true" json:"DB_USERNAME"`
	Password   string `split_words:"true" json:"DB_PASSWORD"`
}

//Db ...
var Db *DbConfig

//loadDbConfig ...
func loadDbConfig() {
	Db = &DbConfig{}
	err := envconfig.Process("db", Db)
	if err != nil {
		log.Fatal(err.Error())
	}
}
