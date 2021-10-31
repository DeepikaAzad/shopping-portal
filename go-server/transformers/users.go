package transformers

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
)

func GetRegisterUserResponse(user entities.Users) models.RegisterUserResp {
	return models.RegisterUserResp{
		Message: user.Name + " registered sucessfully.",
	}
}

func GetLoginUserResponse(token string) models.LoginUserResp {
	return models.LoginUserResp{
		Token: token,
	}
}

func GetUserListResponse(users []entities.Users) models.UserList {
	userList := []models.User{}
	for _, v := range users {
		user := models.User{
			ID:       v.ID,
			UserName: v.UserName,
			FullName: v.Name,
		}
		userList = append(userList, user)
	}
	return models.UserList{
		User: userList,
	}
}
