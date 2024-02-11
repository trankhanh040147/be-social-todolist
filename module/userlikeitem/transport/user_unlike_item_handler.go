package ginuserlikeitem

import (
	"go-200lab-g09/common"
	"go-200lab-g09/module/user/model"
	"go-200lab-g09/module/userlikeitem/biz"
	"go-200lab-g09/module/userlikeitem/storage"
	"net/http"
	"strconv"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func UnlikeItem(serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//id, err := common.FromBase58(ctx.Param("id")) //fake_id
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(*model.User)
		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)

		store := storage.NewSQLStore(db)
		business := biz.NewUserUnlikeItemBiz(store)

		if err := business.UnlikeItem(ctx.Request.Context(), requester.GetUserId(),
			//int(id.GetLocalID())); //fake_id
			id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
