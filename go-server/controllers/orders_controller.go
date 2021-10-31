package controllers

import (
	"net/http"

	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/providers"
	"github.com/DeepikaAzad/go-to-do-app/go-server/transformers"
	"github.com/gin-gonic/gin"
)

func OrderListHandler(ctx *gin.Context) {
	var reqBody models.OrderListReq
	reqBody.UsersID = ctx.GetUint("user_id")
	resp, err := providers.Orders.OrderList(reqBody, ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, transformers.GetOrderListResponse(resp))
}
