package ginitem

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/common"
	"social-todo-list/module/item/biz"
	"social-todo-list/module/item/storage"
)

func GetItem(serviceCtx goservice.ServiceContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := common.UIDFromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// dependency
		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)
		store := storage.NewSQLStore(db)
		business := biz.NewGetItemBiz(store)

		data, err := business.GetItemById(ctx.Request.Context(), int(id.GetLocalID()))

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
