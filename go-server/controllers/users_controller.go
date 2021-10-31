package controllers

import (
	"net/http"

	"github.com/DeepikaAzad/go-to-do-app/go-server/providers"
	"github.com/DeepikaAzad/go-to-do-app/go-server/transformers"
	"github.com/DeepikaAzad/go-to-do-app/go-server/validators"
	"github.com/gin-gonic/gin"
)

// RegisterUserHandler,Register new user
func RegisterUserHandler(ctx *gin.Context) {
	reqBody, szerr := validators.ValidateRegisterUser(ctx)
	if szerr.Errors != nil {
		ctx.Error(&szerr)
		return
	}
	resp, err := providers.Users.RegisterUser(reqBody, ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, transformers.GetRegisterUserResponse(resp))
}

func LoginHandler(ctx *gin.Context) {
	reqBody, szerr := validators.ValidateLoginUser(ctx)
	if szerr.Errors != nil {
		ctx.Error(&szerr)
		return
	}
	resp, err := providers.Users.LoginUser(reqBody, ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, transformers.GetLoginUserResponse(resp.Token))
}

func GetUserListHandler(ctx *gin.Context) {
	resp, err := providers.Users.GetUserList(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, transformers.GetUserListResponse(resp))
}
