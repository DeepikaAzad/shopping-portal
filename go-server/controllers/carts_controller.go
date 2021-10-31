package controllers

import (
	"net/http"

	"github.com/DeepikaAzad/go-to-do-app/go-server/constants"
	"github.com/DeepikaAzad/go-to-do-app/go-server/providers"
	"github.com/DeepikaAzad/go-to-do-app/go-server/transformers"
	"github.com/DeepikaAzad/go-to-do-app/go-server/validators"
	"github.com/gin-gonic/gin"
)

func AddItemToCartHandler(ctx *gin.Context) {
	reqBody, slErr := validators.ValidateAddItemToCart(ctx)
	if slErr.Errors != nil {
		ctx.Error(&slErr)
		return
	}
	resp, err := providers.Carts.AddItemToCart(reqBody, ctx)
	if err != nil {
		if err.Error() == constants.ErrorCode.DUPLICATE_ERROR {
			ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		}
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, transformers.GetAddItemToCartResponse(resp))
}

func PlaceOrderHandler(ctx *gin.Context) {
	reqBody, slErr := validators.ValidatePlaceOrderHandler(ctx)
	if slErr.Errors != nil {
		ctx.Error(&slErr)
		return
	}
	resp, err := providers.Carts.PlaceOrder(reqBody, ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, transformers.GetPlaceOrderResponse(resp))
}
