package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/DeepikaAzad/go-to-do-app/go-server/constants"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func AppErrorReporter() gin.HandlerFunc {
	return jsonAppErrorReporterT(gin.ErrorTypeAny)
}

func jsonAppErrorReporterT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if errPanicCtx := recover(); errPanicCtx != nil {
				var parsedErrorFormatted models.SZLError
				parsedErrorFormatted.Message = "something went wrong"
				parsedErrorFormatted.Type = constants.ErrorCode.INTERNAL_SERVER_ERROR
				setFormattedErrorInJson(parsedErrorFormatted, c)
				return
			}
		}()

		c.Next()
		detectedErrors := c.Errors.ByType(errType)
		var parsedErrorFormatted models.SZLError
		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			switch errors.Cause(err).(type) {
			case *models.SZLError:
				parsedError := errors.Cause(err).(*models.SZLError)
				parsedErrorFormatted.Type = parsedError.Type
				parsedErrorFormatted.Message = parsedError.Message
			default:
				parsedErrorFormatted.Message = err.Error()
			}

			setFormattedErrorInJson(parsedErrorFormatted, c)
		}
	}
}

func setFormattedErrorInJson(parsedErrorFormatted models.SZLError, c *gin.Context) {
	code := getHTTPstatusCode(parsedErrorFormatted.Type)
	tempbuff := new(bytes.Buffer)
	json.NewEncoder(tempbuff).Encode(parsedErrorFormatted)
	fr := tempbuff.String()
	c.Set("statusCode", code)
	c.Set("response", fr)
	c.JSON(code, parsedErrorFormatted)
	c.Abort()

}

func getHTTPstatusCode(message string) int {
	if message == constants.ErrorCode.INTERNAL_SERVER_ERROR {
		return http.StatusInternalServerError
	}
	if message == constants.ErrorCode.INVALID_PASSWORD {
		return http.StatusUnauthorized
	}
	if message == constants.ErrorCode.INVALID_TOKEN {
		return http.StatusUnauthorized
	}
	if message == constants.ErrorCode.INVALID_REQUEST_ERROR || message == constants.ErrorCode.DUPLICATE_ERROR {
		return http.StatusUnprocessableEntity
	}
	if message == constants.ErrorCode.NOT_FOUND {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
