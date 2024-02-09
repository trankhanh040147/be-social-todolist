package ginitem

import (
	"go-200lab-g09/common"
	"go-200lab-g09/module/item/biz"
	"go-200lab-g09/module/item/model"
	"go-200lab-g09/module/item/storage"
	"net/http"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(serviceCtx goservice.ServiceContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var itemData model.TodoItemCreation

		if err := c.ShouldBind(&itemData); err != nil {
			// >> c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),}) is called. This sends a JSON response to the client with a status code of 400 (Bad Request). The JSON body of the response contains a single property error, which is set to the string representation of the error returned by ShouldBind. The gin.H function is a shortcut for creating a map in Go, which in this case is used to create the JSON object.
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		itemData.UserId = requester.GetUserId()
		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)

		store := storage.NewSQLStore(db)
		business := biz.NewCreateItemBiz(store)

		// step 3: use db.Create to
		if err := business.CreateNewItem(c.Request.Context(), &itemData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// step 4: print data of the inserted record
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(itemData.Id))
	}
}
