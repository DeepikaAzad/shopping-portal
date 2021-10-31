package transformers

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
)

func GetRegisterUserResponse(user entities.Users) models.RegisterUserResp {
	return models.RegisterUserResp{
		Message: user.Name + " registered suucessfully.",
	}
}

func GetLoginUserResponse(token string) models.LoginUserResp {
	return models.LoginUserResp{
		Token: token,
	}
}
