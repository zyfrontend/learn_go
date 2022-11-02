package router

import (
	"App/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func TestMiddleWare(r *gin.Engine) {
	// 局部中间件
	r.GET("/ce", middleware.MiddleWare(), func(c *gin.Context) {
		req, _ := c.Get("request")
		fmt.Println("request:", req)
		// 接收
		c.JSON(200, gin.H{"request": req})
	})
}
