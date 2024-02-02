package ginitem

import (
	"go-200lab-g09/common"
	"go-200lab-g09/module/item/biz"
	"go-200lab-g09/module/item/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// fmt.Println([]int{}[0])

		// dependency
		store := storage.NewSQLStore(db)
		business := biz.NewGetItemBiz(store)

		data, err := business.GetItemById(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
