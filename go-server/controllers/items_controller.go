package controllers

import (
	"net/http"

	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/providers"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func AddItemHandler(ctx *gin.Context) {
	// @TODO:: Validations
	reqBody := models.AddItemReq{}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.Error(errors.Wrap(err, "[AddItemHandler]"))
		return
	}
	resp, err := providers.Items.AddItem(reqBody, ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"name": resp.Name,
	})
}
