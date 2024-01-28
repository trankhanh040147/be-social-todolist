package ginitem

import (
	"go-200lab-g09/common"
	"go-200lab-g09/module/item/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Process()

		var result []model.TodoItem

		db = db.Table(model.TodoItem{}.TableName()).Where("status <> ?", "Deleted")

		// SELECT COUNT(id) FROM todo_items --> paging.Total
		if err := db.Table(model.TodoItem{}.TableName()).Select("id").Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		// SELECT * FROM todo_items ORDER BY id ASC LIMIT paging.Limit OFFSET (paging.Page - 1) * paging.Limit
		if err := db.Table(model.TodoItem{}.TableName()).
			Select("*").
			Offset((paging.Page - 1) * paging.Limit).
			Limit(paging.Limit).
			Order("id asc").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
