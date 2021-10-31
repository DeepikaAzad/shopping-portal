package items

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemsGorm struct{}

func (r ItemsGorm) CreateItem(item entities.Items, ctx *gin.Context) (entities.Items, error) {
	dbx, _ := ctx.Get("DB")
	db := dbx.(*gorm.DB)
	err := db.Create(&item).Error
	if err != nil {
		return item, err
	}
	return item, nil
}

func (r ItemsGorm) GetItemList(ctx *gin.Context) ([]entities.Items, error) {
	dbx, _ := ctx.Get("DB")
	db := dbx.(*gorm.DB)
	var items []entities.Items
	rows, err := db.Find(&items).Rows()
	if err != nil {
		return items, err
	}
	defer rows.Close()
	return items, nil
}

func (r ItemsGorm) GetItemByName(name string, ctx *gin.Context) (entities.Items, error) {
	dbx, _ := ctx.Get("DB")
	db := dbx.(*gorm.DB)
	var item entities.Items
	err := db.Where("name", name).Take(&item).Error
	if err != nil {
		return item, err
	}
	return item, nil
}
