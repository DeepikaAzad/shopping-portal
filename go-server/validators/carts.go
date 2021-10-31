package validators

import (
	"errors"

	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/providers/repositories"
	"github.com/DeepikaAzad/go-to-do-app/go-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gorm.io/gorm"
)

func ValidateAddItemToCart(ctx *gin.Context) (models.AddItemToCartReq, models.SZLError) {
	var reqBody models.AddItemToCartReq
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return reqBody, GetInvalidDataTypeSzlError(err)
	}

	rules := govalidator.MapData{
		"item_name": []string{"between:1,15", "alpha"},
		"item_id":   []string{"required", "numeric"},
	}

	opts := govalidator.Options{
		Data:  &reqBody, // request object
		Rules: rules,    // rules map
	}
	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		err := GetValidationError(e)
		return reqBody, err
	}

	item, err := repositories.Items.GetItemByID(reqBody.ItemID, ctx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return reqBody, GetInternalServerError(err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return reqBody, GetNotFoundError(errors.New("item not found"))
	}
	reqBody.ItemID = uint(item.ID)
	return reqBody, models.SZLError{}
}

func ValidatePlaceOrderHandler(ctx *gin.Context) (models.PlaceOrderReq, models.SZLError) {
	var reqBody models.PlaceOrderReq
	reqBody.CartID = utils.ParseUint(ctx.Param("cartId"))
	reqBody.UserID = ctx.GetUint("user_id")
	cart, err := repositories.Carts.GetCartByIdAndUserId(reqBody.UserID, reqBody.CartID, ctx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return reqBody, GetInternalServerError(err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return reqBody, GetNotFoundError(errors.New("cart does not exist"))
	}
	if cart.IsPurchased == 1 {
		return reqBody, GetNotFoundError(errors.New("cart is empty"))
	}
	return reqBody, models.SZLError{}
}
