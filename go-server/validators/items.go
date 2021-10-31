package validators

import (
	"errors"

	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/providers/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ValidateAddItem(ctx *gin.Context) (models.AddItemReq, models.SZLError) {
	var reqBody models.AddItemReq
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return reqBody, GetInvalidDataTypeSzlError(err)
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
