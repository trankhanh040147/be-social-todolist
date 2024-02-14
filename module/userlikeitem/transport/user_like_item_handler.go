package ginuserlikeitem

import (
	"net/http"
	"social-todo-list/common"
	itemStorage "social-todo-list/module/item/storage"
	"social-todo-list/module/userlikeitem/biz"
	"social-todo-list/module/userlikeitem/model"
	"social-todo-list/module/userlikeitem/storage"
	"time"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LikeItem(serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//id, err := strconv.Atoi(ctx.Param("id"))
		id, err := common.UIDFromBase58(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)

		store := storage.NewSQLStore(db)
		itemStore := itemStorage.NewSQLStore(db)
		business := biz.NewUserLikeItemBiz(store, itemStore)
		now := time.Now().UTC()

		if err := business.LikeItem(ctx.Request.Context(), &model.Like{
			UserId: requester.GetUserId(),
			ItemId: int(id.GetLocalID()),
			//ItemId:    id,
			CreatedAt: &now,
		}); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
