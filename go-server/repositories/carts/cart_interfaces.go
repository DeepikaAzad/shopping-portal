package carts

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
	"github.com/gin-gonic/gin"
)

type CartGormInterface interface {
	CreateCart(entities.Carts, *gin.Context) (entities.Carts, error)
	GetCartByUserAndPurchasedFalse(uint, *gin.Context) (entities.Carts, error)
	UpdateCartItemIDs(entities.Carts, uint, *gin.Context) (entities.Carts, error)
	DeleteCart(entities.Carts, *gin.Context) error
	UpdateCartIsPurchased(int8, uint, *gin.Context) (entities.Carts, error)
	GetCartByIdAndUserId(uint, uint, *gin.Context) (entities.Carts, error)
}
