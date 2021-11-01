package controllers

import (
	"errors"
	"net/http"

	"github.com/DeepikaAzad/go-to-do-app/go-server/constants"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
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
			ctx.JSON(http.StatusUnprocessableEntity, validators.GetDuplicateError(errors.New("item already exist")))
			return
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

func GetCartHandler(ctx *gin.Context) {
	reqBody := models.CartListReq{}
	reqBody.UserID = ctx.GetUint("user_id")
	resp, err := providers.Carts.GetCart(reqBody, ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, transformers.GetCartResponse(resp))
}

func RemoveItemFromCartHandler(ctx *gin.Context) {
	reqBody, slErr := validators.ValidateRemoveItemFromCart(ctx)
	if slErr.Errors != nil {
		ctx.Error(&slErr)
		return
	}
	err := providers.Carts.RemoveItemFromCart(reqBody, ctx)
	if err != nil {
		if err.Error() == constants.ErrorCode.NOT_FOUND {
			ctx.JSON(http.StatusBadRequest, validators.GetNotFoundError(errors.New("no cart found")))
			return
		}
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, transformers.GetRemoveItemFromCartResponse())
}
