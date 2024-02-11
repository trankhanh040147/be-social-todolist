package ginuserlikeitem

import (
	"go-200lab-g09/common"
	"go-200lab-g09/module/userlikeitem/biz"
	"go-200lab-g09/module/userlikeitem/model"
	"go-200lab-g09/module/userlikeitem/storage"
	"net/http"
	"strconv"
	"time"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LikeItem(serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		// id, err := common.FromBase58(ctx.Param("id")) // fake_id
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)

		store := storage.NewSQLStore(db)
		business := biz.NewUserLikeItemBiz(store)
		now := time.Now().UTC()

		if err := business.LikeItem(ctx.Request.Context(), &model.Like{
			UserId: requester.GetUserId(),
			// fake_id
			// ItemId:    int(id.GetLocalID()),
			ItemId:    id,
			CreatedAt: &now,
		}); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
