package ginitem

import (
	"go-200lab-g09/common"
	"go-200lab-g09/module/item/biz"
	"go-200lab-g09/module/item/model"
	"go-200lab-g09/module/item/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			common.Paging
			model.Filter
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		queryString.Paging.Process()

		store := storage.NewSQLStore(db)
		business := biz.NewListItemBiz(store)

		result, err := business.ListItem(c.Request.Context(), &queryString.Filter, &queryString.Paging)

		// SELECT * FROM todo_items ORDER BY id ASC LIMIT paging.Limit OFFSET (paging.Page - 1) * paging.Limit
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, queryString.Paging, queryString.Filter))
	}
}
