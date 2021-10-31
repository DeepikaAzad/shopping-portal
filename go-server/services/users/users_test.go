package users_test

import (
	"github.com/DeepikaAzad/go-to-do-app/go-server/app"
	"github.com/DeepikaAzad/go-to-do-app/go-server/models"
	"github.com/DeepikaAzad/go-to-do-app/go-server/services/users"
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
}

var _ = Describe("Users", func() {
	Describe("RegisterUser", func() {
		It("Register User", func() {
			user := models.RegisterUserReq{
				UserName: "test",
				FullName: "test",
				Password: "dGVzdEAxMjM=",
			}
			resp, _ := users.UsersImpl{}.RegisterUser(user, ctx)
			gomega.Expect(resp.UserName).To(gomega.Equal("test"))
		})
	})
})
