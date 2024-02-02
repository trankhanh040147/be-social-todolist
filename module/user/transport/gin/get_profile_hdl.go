package ginuser

import (
	"go-200lab-g09/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(user))

	}
}
