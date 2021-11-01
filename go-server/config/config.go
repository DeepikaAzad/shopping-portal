package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

//LoadConfigs ...
func LoadConfigs() {

	// Get current file full path from runtime
	_, b, _, _ := runtime.Caller(0)

	// Root folder of this project
	ProjectRootPath := filepath.Join(filepath.Dir(b), "../")
	// load .env file
	err := godotenv.Load(ProjectRootPath + "/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	loadAppConfig()
	loadDbConfig()
}
