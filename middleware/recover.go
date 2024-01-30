package middleware

import (
	"go-200lab-g09/common"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json;")

				// case: Error from business logic, or db
				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					// panic(err)
					return
				}

				// case: Not AppErr, just root error.
				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err) // show stack trace
				return
			}
		}()
		c.Next()
	}
}
