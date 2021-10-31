package carts

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
	"github.com/DeepikaAzad/go-to-do-app/go-server/providers/repositories"
	"github.com/DeepikaAzad/go-to-do-app/go-server/services"
	"github.com/DeepikaAzad/go-to-do-app/go-server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartsImpl struct{}

func (i CartsImpl) PlaceOrder(req models.PlaceOrderReq, ctx *gin.Context) (entities.Carts, error) {

	cart := entities.Carts{
		IsPurchased: 1,
	}
	// update cart
	cart, err := repositories.Carts.UpdateCartIsPurchased(cart.IsPurchased, req.CartID, ctx)
	if err != nil {
		return cart, err
	}
	// save order
	order := entities.Orders{
		CartsID: req.CartID,
		UsersID: req.UserID,
	}
	_, err = repositories.Orders.CreateOrder(order, ctx)
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (i CartsImpl) AddItemToCart(req models.AddItemToCartReq, ctx *gin.Context) (entities.Carts, error) {
	userID := ctx.GetUint("user_id")
	// Check if cart exist for user
	cart, err := repositories.Carts.GetCartByUserAndPurchasedFalse(userID, ctx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return cart, err
	}
	itemsId, itemErr := addItem(cart.ItemsID, req.ItemID)
	if itemErr != nil {
		return cart, itemErr
	}

	cart = entities.Carts{
		UsersID:     userID,
		IsPurchased: 0,
		ItemsID:     itemsId,
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Add cart
		cart, err := repositories.Carts.CreateCart(cart, ctx)
		if err != nil {
			return cart, err
		}
	} else {
		// YES: Update with Item name and check dulicate items
		cart, err := repositories.Carts.UpdateCart(cart, cart.ID, ctx)
		if err != nil {
			return cart, err
		}
	}

	return cart, nil
}

func addItem(itemsID string, itemId uint) (string, error) {
	itemJson := []string{}

	if itemsID == "" {
		itemStr, _ := json.Marshal(itemJson)
		err := json.Unmarshal(itemStr, &itemJson)
		if err != nil {
			return "", err
		}
	} else {
		err := json.Unmarshal([]byte(itemsID), &itemJson)
		if err != nil {
			return "", err
		}
	}

	// Check duplicate item
	for _, id := range itemJson {
		if utils.ParseUint(id) == itemId {
			return "", errors.New("duplicate_item")
		}
	}

	itemJson = append(itemJson, fmt.Sprint(itemId))

	resp, _ := json.Marshal(itemJson)
	return string(resp), nil
}

func (i CartsImpl) GetCart(req models.CartListReq, ctx *gin.Context) (models.CartListResp, error) {
	resp := models.CartListResp{}
	cart, err := repositories.Carts.GetCartByUserAndPurchasedFalse(req.UserID, ctx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return resp, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return resp, err
	}

	itemNameList, err := services.AddNameInCardItemIds(cart.ItemsID, ctx)
	if err != nil {
		return resp, err
	}
	resp.CartID = cart.ID
	resp.Items = itemNameList
	return resp, nil
}
