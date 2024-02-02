package ginuser

import (
	"go-200lab-g09/common"
	"go-200lab-g09/component/tokenprovider"
	"go-200lab-g09/module/user/biz"
	"go-200lab-g09/module/user/model"
	"go-200lab-g09/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB, tokenProvider tokenprovider.Provider) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData model.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		md5 := common.NewMd5Hash()

		business := biz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*7)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))

	}
}
