package validators

import (
	"errors"

	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/providers/repositories"
	"github.com/DeepikaAzad/go-to-do-app/go-server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ValidateRegisterUser(ctx *gin.Context) (models.RegisterUserReq, models.SZLError) {
	var reqBody models.RegisterUserReq
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return reqBody, GetInvalidDataTypeSzlError(err)
	}
	user, err := repositories.Users.GetUserByUserName(reqBody.UserName, ctx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return reqBody, GetInternalServerError(err)
	}

	if user.ID != 0 {
		return reqBody, GetDuplicateError(errors.New("user name already exist"))
	}

	return reqBody, models.SZLError{}
}

func ValidateLoginUser(ctx *gin.Context) (models.LoginUserReq, models.SZLError) {
	var reqBody models.LoginUserReq
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return reqBody, GetInvalidDataTypeSzlError(err)
	}

	// User exist
	user, err := repositories.Users.GetUserByUserName(reqBody.UserName, ctx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return reqBody, GetInternalServerError(err)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return reqBody, GetNotFoundError(errors.New("user name not found"))
	}

	// Match password
	reqBody.Password = utils.DecodeBase64(reqBody.Password)
	if !utils.CheckPasswordHash(reqBody.Password, user.Password) {
		return reqBody, GetInvalidPwdError(errors.New("invalid password"))
	}
	return reqBody, models.SZLError{}
}
