package controllers

import (
	"net/http"

	"github.com/DeepikaAzad/go-to-do-app/go-server/providers"
	"github.com/DeepikaAzad/go-to-do-app/go-server/transformers"
	"github.com/DeepikaAzad/go-to-do-app/go-server/validators"
	"github.com/gin-gonic/gin"
)

func AddItemHandler(ctx *gin.Context) {
	reqBody, slErr := validators.ValidateAddItem(ctx)
	if slErr.Errors != nil {
		ctx.Error(&slErr)
		return
	}
	resp, err := providers.Items.AddItem(reqBody, ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, transformers.GetAddItemResponse(resp))
}
