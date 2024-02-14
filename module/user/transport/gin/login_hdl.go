package ginuser

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/module/user/biz"
	"social-todo-list/module/user/model"
	"social-todo-list/module/user/storage"
	"social-todo-list/plugin/tokenprovider"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData model.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)
		tokenProvider := serviceCtx.MustGet(common.PluginJWT).(tokenprovider.TokenProvider)
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
