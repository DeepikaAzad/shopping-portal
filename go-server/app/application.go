package app

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/DeepikaAzad/go-to-do-app/go-server/app/middleware"
	"github.com/DeepikaAzad/go-to-do-app/go-server/config"
	"github.com/DeepikaAzad/go-to-do-app/go-server/database"
	"github.com/DeepikaAzad/go-to-do-app/go-server/database/commands"
	"github.com/spf13/cobra"
	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginglog "github.com/szuecs/gin-glog"
)

var (
	router = gin.Default()
	DB     *gorm.DB
)

func StartApplication() {
	Init()
	router.Use(ginglog.Logger(3 * time.Second))
	fmt.Println("Rest API - Gin Gonic")
	registerMiddlewares()
	Router()
	router.Run(":" + strconv.Itoa(config.App.Port))
}

func Init() {
	config.LoadConfigs()
	initDB()
}

func initDB() {
	db, _ := database.Connection()
	DB = db
	cmd := &cobra.Command{}
	cmd.AddCommand(commands.Migrate())
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}

func registerMiddlewares() {
	router.Use(middleware.AppErrorReporter())
	router.Use(middleware.DatabaseContext(DB))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://127.0.0.1:3000", "http://localhost:3000"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("*")

	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	corsConfig.AddAllowHeaders(allowedHeaders)
	router.Use(cors.New(corsConfig))
}
