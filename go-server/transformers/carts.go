package transformers

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
)

func GetAddItemToCartResponse(cart entities.Carts) models.AddItemToCartResp {
	return models.AddItemToCartResp{
		Message: "item added sucessfully to your cart",
	}
}

func GetRemoveItemFromCartResponse() models.RemoveItemfromCartResp {
	return models.RemoveItemfromCartResp{
		Message: "item removed sucessfully from your cart",
	}
}

func GetPlaceOrderResponse(cart entities.Carts) models.PlaceOrderResp {
	return models.PlaceOrderResp{
		CartID:  cart.ID,
		Message: "item orderd sucessfully",
	}
}

func GetCartResponse(cart models.CartListResp) models.CartListResp {
	if cart.CartID == 0 {
		cart.Items = []models.Item{}
	}
	return cart
}
