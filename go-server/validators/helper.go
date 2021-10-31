package validators

import (
	"errors"

	"github.com/DeepikaAzad/go-to-do-app/go-server/constants"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
)

func GetInvalidDataTypeSzlError(err error) models.SZLError {
	msg := "request body is invalid"
	szErr := models.SZLError{
		Message: msg,
		Type:    constants.ErrorCode.INVALID_REQUEST_ERROR,
		Errors:  errors.New(msg),
	}
	return szErr
}

func GetInternalServerError(err error) models.SZLError {
	return models.SZLError{
		Type:    constants.ErrorCode.INTERNAL_SERVER_ERROR,
		Message: err.Error(),
		Errors:  err,
	}
}

func GetInvalidRequestError(err error) models.SZLError {
	return models.SZLError{
		Type:    constants.ErrorCode.INTERNAL_SERVER_ERROR,
		Message: err.Error(),
		Errors:  err,
	}
}

func GetDuplicateError(err error) models.SZLError {
	return models.SZLError{
		Message: err.Error(),
		Type:    constants.ErrorCode.DUPLICATE_ERROR,
		Errors:  err,
	}
}

func GetNotFoundError(err error) models.SZLError {
	return models.SZLError{
		Message: err.Error(),
		Type:    constants.ErrorCode.NOT_FOUND,
		Errors:  err,
	}
}

func GetInvalidPwdError(err error) models.SZLError {
	return models.SZLError{
		Message: err.Error(),
		Type:    constants.ErrorCode.INVALID_PASSWORD,
		Errors:  err,
	}
}

func GetInvalidTokenError(err error) models.SZLError {
	return models.SZLError{
		Message: err.Error(),
		Type:    constants.ErrorCode.INVALID_TOKEN,
		Errors:  err,
	}
}
