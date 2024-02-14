package ginitem

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/module/item/biz"
	"social-todo-list/module/item/storage"
	"strconv"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItem(serviceCtx goservice.ServiceContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// dependency
		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)
		store := storage.NewSQLStore(db)
		business := biz.NewGetItemBiz(store)

		data, err := business.GetItemById(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
