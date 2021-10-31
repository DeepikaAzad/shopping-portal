package users

import (
	"fmt"

	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UsersGorm struct{}

func (r UsersGorm) CreateUser(user entities.Users, ctx *gin.Context) (entities.Users, error) {
	dbx, f := ctx.Get("DB")
	if !f {
		// Handle error
		fmt.Println(f)
	}

	db := dbx.(*gorm.DB)
	err := db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r UsersGorm) UpdateUser(user entities.Users, userName string, ctx *gin.Context) (entities.Users, error) {
	dbx, _ := ctx.Get("DB")
	db := dbx.(*gorm.DB)
	err := db.Model(&entities.Users{}).Where("user_name", userName).Updates(user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r UsersGorm) GetUserByUserName(userName string, ctx *gin.Context) (entities.Users, error) {
	dbx, f := ctx.Get("DB")
	if !f {
		// Handle error
		fmt.Println(f)
	}
	db := dbx.(*gorm.DB)
	var user entities.Users
	err := db.Where("user_name", userName).Take(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r UsersGorm) GetUserList(ctx *gin.Context) ([]entities.Users, error) {
	dbx, _ := ctx.Get("DB")
	db := dbx.(*gorm.DB)
	var users []entities.Users
	rows, err := db.Find(&users).Rows()
	if err != nil {
		return users, err
	}
	defer rows.Close()
	return users, nil
}
