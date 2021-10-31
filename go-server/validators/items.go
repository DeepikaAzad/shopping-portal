package validators

import (
	"errors"

	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/providers/repositories"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gorm.io/gorm"
)

func ValidateAddItem(ctx *gin.Context) (models.AddItemReq, models.SZLError) {
	var reqBody models.AddItemReq
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return reqBody, GetInvalidDataTypeSzlError(err)
	}

	rules := govalidator.MapData{
		"name": []string{"required", "between:1,15", "alpha_space"},
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

	_, err := repositories.Items.GetItemByName(reqBody.Name, ctx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return reqBody, GetInternalServerError(err)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return reqBody, GetDuplicateError(errors.New("item already exist"))
	}

	return reqBody, models.SZLError{}
}
