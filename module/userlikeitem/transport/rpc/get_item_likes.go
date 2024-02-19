package rpcuserlikeitem

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/module/userlikeitem/storage"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItemLikes(serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type requestData struct {
			Ids []int `json:"ids"`
		}

		var data requestData
		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)
		store := storage.NewSQLStore(db)

		mapResult, err := store.GetItemLikes(ctx.Request.Context(), data.Ids)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(mapResult))
	}
}
