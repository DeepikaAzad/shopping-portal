package controllers

import (
	"net/http"

	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/providers"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// RegisterUserHandler,Register new user
func RegisterUserHandler(ctx *gin.Context) {
	// @TODO:: Validations
	reqBody := models.RegisterUserReq{}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.Error(errors.Wrap(err, "[RegisterUserHandler]"))
		return
	}
	resp, err := providers.Users.RegisterUser(reqBody, ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":        resp.ID,
		"name":      resp.Name,
		"user_name": resp.UserName,
		"token":     resp.Token,
	})
}

func LoginHandler(ctx *gin.Context) {
	// @TODO:: Validations
	reqBody := models.LoginUserReq{}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.Error(errors.Wrap(err, "[RegisterUserHandler]"))
		return
	}
	resp, err := providers.Users.LoginUser(reqBody, ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": resp.Token,
	})
}
