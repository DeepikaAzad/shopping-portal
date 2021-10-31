package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// AppConfig ..
type AppConfig struct {
	Debug bool
	Port  int
	Env   string
}

//App ..
var App *AppConfig

func loadAppConfig() {
	App = &AppConfig{}
	// var s AppConfig
	err := envconfig.Process("app", App)
	if err != nil {
		log.Fatal(err.Error())
	}
}
