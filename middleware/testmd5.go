package middleware

import (
	"app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
)

func TestMd5() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(405, gin.H{
				"code":    405,
				"message": "未登录或非法访问",
			})
			c.Abort()
			return
		}
		claims, err := tools.VerifyToken(token)
		fmt.Println("claims, err", claims, err)
		if err != nil {
			c.JSON(405, gin.H{
				"code":    405,
				"message": "您的帐户异地登陆或令牌失效",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
