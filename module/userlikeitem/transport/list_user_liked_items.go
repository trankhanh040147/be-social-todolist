package ginuserlikeitem

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"go-200lab-g09/common"
	"go-200lab-g09/module/userlikeitem/biz"
	"go-200lab-g09/module/userlikeitem/storage"
	"net/http"

	"gorm.io/gorm"
)

func ListLikedUsers(serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := common.UIDFromBase58(ctx.Param("id"))
		//id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Process()

		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)

		store := storage.NewSQLStore(db)
		business := biz.NewListUsersLikedItemBiz(store)

		result, err := business.ListUsersLikedItem(ctx.Request.Context(),
			int(id.GetLocalID()),
			&paging)
		if err != nil {
			panic(err)
		}

		for index := range result {
			result[index].Mask()
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
