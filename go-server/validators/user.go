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

func ValidateRegisterUser(ctx *gin.Context) (models.RegisterUserReq, models.SZLError) {
	var reqBody models.RegisterUserReq
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return reqBody, GetInvalidDataTypeSzlError(err)
	}

	rules := govalidator.MapData{
		"user_name": []string{"required", "between:3,15"},
		"full_name": []string{"required"},
		"password":  []string{"required"},
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

	user, err := repositories.Users.GetUserByUserName(reqBody.UserName, ctx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return reqBody, GetInternalServerError(err)
	}

	if user.ID != 0 {
		return reqBody, GetDuplicateError(errors.New("user name already exist"))
	}

	return reqBody, models.SZLError{}
}

func ValidateLoginUser(ctx *gin.Context) (reqBody models.LoginUserReq, szErr models.SZLError) {
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return reqBody, GetInvalidDataTypeSzlError(err)
	}
	rules := govalidator.MapData{
		"user_name": []string{"required"},
		"password":  []string{"required"},
	}
	opts := govalidator.Options{
		Data:            &reqBody, // request object
		Rules:           rules,    // rules map
		RequiredDefault: true,     // all the field to be pass the rules
	}
	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		err := GetValidationError(e)
		return reqBody, err
	}

	// User exist
	user, err := repositories.Users.GetUserByUserName(reqBody.UserName, ctx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return reqBody, GetInternalServerError(err)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return reqBody, GetInvalidUserError(errors.New("invalid user"))
	}

	// Match password
	reqBody.Password = utils.DecodeBase64(reqBody.Password)
	if !utils.CheckPasswordHash(reqBody.Password, user.Password) {
		return reqBody, GetInvalidUserError(errors.New("invalid password"))
	}
	return reqBody, models.SZLError{}
}
