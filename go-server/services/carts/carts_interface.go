package carts

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
	"github.com/gin-gonic/gin"
)

type CartsInterface interface {
	AddItemToCart(models.AddItemToCartReq, *gin.Context) (entities.Carts, error)
	PlaceOrder(models.PlaceOrderReq, *gin.Context) (entities.Carts, error)
	GetCart(models.CartListReq, *gin.Context) (models.CartListResp, error)
}
