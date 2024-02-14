package ginuserlikeitem

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"net/http"
	"social-todo-list/common"
	itemStorage "social-todo-list/module/item/storage"
	"social-todo-list/module/user/model"
	"social-todo-list/module/userlikeitem/biz"
	"social-todo-list/module/userlikeitem/storage"

	"gorm.io/gorm"
)

func UnlikeItem(serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := common.UIDFromBase58(ctx.Param("id")) //fake_id
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(*model.User)
		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)

		store := storage.NewSQLStore(db)
		itemStore := itemStorage.NewSQLStore(db)
		business := biz.NewUserUnlikeItemBiz(store, itemStore)

		if err := business.UnlikeItem(ctx.Request.Context(), requester.GetUserId(),
			int(id.GetLocalID()),
		); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
