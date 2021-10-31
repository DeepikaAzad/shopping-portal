package items_test

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/app"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models/entities"
	"github.com/DeepikaAzad/go-to-do-app/go-server/services/items"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

var (
	ctx *gin.Context
	db  *gorm.DB
)

func init() {
	ctx = app.Bootstrap()
	dbx, _ := ctx.Get("DB")
	db = dbx.(*gorm.DB)
}

var _ = Describe("Items", func() {
	var resp entities.Items
	AfterEach(func() {
		db.Delete(&resp)
	})
	Describe("AddItem", func() {
		It("Add Item", func() {
			item := models.AddItemReq{
				Name: "test Item",
			}
			resp, _ = items.ItemsImpl{}.AddItem(item, ctx)
			gomega.Expect(resp.Name).To(gomega.Equal("test Item"))
		})
	})
})
