package carts

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartsGorm struct{}

func (r CartsGorm) CreateCart(cart entities.Carts, ctx *gin.Context) (entities.Carts, error) {
	dbx, _ := ctx.Get("DB")
	db := dbx.(*gorm.DB)
	err := db.Create(&cart).Error
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (r CartsGorm) GetCartByUserAndPurchasedFalse(userId uint, ctx *gin.Context) (entities.Carts, error) {
	dbx, _ := ctx.Get("DB")
	db := dbx.(*gorm.DB)
	var cart entities.Carts
	err := db.Where("users_id = ? and is_purchased = ? ", userId, 0).Take(&cart).Error
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (r CartsGorm) UpdateCartItemIDs(cart entities.Carts, cartId uint, ctx *gin.Context) (entities.Carts, error) {
	dbx, _ := ctx.Get("DB")
	db := dbx.(*gorm.DB)
	err := db.Model(&cart).Where("id", cart.ID).Update("items_id", cart.ItemsID).Error
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (r CartsGorm) DeleteCart(cart entities.Carts, ctx *gin.Context) error {
	dbx, _ := ctx.Get("DB")
	db := dbx.(*gorm.DB)
	err := db.Where("id", cart.ID).Delete(&cart).Error
	if err != nil {
		return err
	}
	return nil
}

func (r CartsGorm) UpdateCartIsPurchased(isPurchansed int8, cartId uint, ctx *gin.Context) (entities.Carts, error) {
	dbx, _ := ctx.Get("DB")
	db := dbx.(*gorm.DB)
	cart := entities.Carts{}
	err := db.Model(&cart).Where("id", cartId).Update("is_purchased", isPurchansed).Error
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (r CartsGorm) GetCartByIdAndUserId(userId, cartId uint, ctx *gin.Context) (entities.Carts, error) {
	dbx, _ := ctx.Get("DB")
	db := dbx.(*gorm.DB)
	var cart entities.Carts
	err := db.Where("users_id = ? and id = ?", userId, cartId, 0).Take(&cart).Error
	if err != nil {
		return cart, err
	}
	return cart, nil
}
