package middleware

import (
	"app/modules/common/response"
	"app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
)

func VerifyPermissions() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.Fail(c)
			c.Abort()
			return
		}
		_, err := tools.VerifyToken(token)
		fmt.Println("err", err)
		if err != nil {
			response.Fail(c)
			c.Abort()
			return
		}
		c.Next()
	}
}
