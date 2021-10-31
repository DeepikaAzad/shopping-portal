package transformers

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
)

func GetAddItemToCartResponse(cart entities.Carts) models.AddItemToCartResp {
	return models.AddItemToCartResp{
		CartID:  cart.ID,
		ItemsID: cart.ItemsID,
		Message: "item added sucessfully to your cart",
	}
}

func GetPlaceOrderResponse(cart entities.Carts) models.PlaceOrderResp {
	return models.PlaceOrderResp{
		CartID:  cart.ID,
		Message: "item orderd sucessfully",
	}
}

func GetCartResponse(cart models.CartListResp) models.CartListResp {
	return cart
}
