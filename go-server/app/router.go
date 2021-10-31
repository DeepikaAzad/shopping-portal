package app

import (
	"net/http"

	"github.com/DeepikaAzad/go-to-do-app/go-server/app/middleware"
	"github.com/DeepikaAzad/go-to-do-app/go-server/controllers"
	"github.com/gin-gonic/gin"
)

// Router is exported and used in main.go
func Router() {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	api := router.Group("/api")
	user := api.Group("/user")
	user.POST("/create", controllers.RegisterUserHandler)
	user.POST("/login", controllers.LoginHandler)

	api.Use(middleware.AuthorizeJWT())
	api.POST("/item/create", controllers.AddItemHandler)

}
