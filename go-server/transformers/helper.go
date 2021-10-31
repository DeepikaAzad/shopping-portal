package transformers

import "github.com/DeepikaAzad/go-to-do-app/go-server/models"

func ErrorResponse(msg string) models.ErrorResp {
	return models.ErrorResp{
		Message: msg,
	}
}
