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
	user.GET("/list", controllers.GetUserListHandler)

	api.Use(middleware.AuthorizeJWT())
	api.POST("/item/create", controllers.AddItemHandler)
	api.GET("/item/list", controllers.GetItemListHandler)
	api.POST("/cart/add", controllers.AddItemToCartHandler)
	api.POST("/cart/:cartId/complete", controllers.PlaceOrderHandler)
	api.GET("/cart/list", controllers.GetCartHandler)
	api.GET("/order/list", controllers.OrderListHandler)

}
