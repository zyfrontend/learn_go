package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Get(r *gin.Engine) {
	// 表单参数
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("userName")
		password := c.PostForm("userPassword")
		c.String(http.StatusOK, fmt.Sprintf("username: %s, password: %s, type: %s", username, password, types))
	})
	req := func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	}
	r.GET("/", req)
	r.POST("/xxxpost", req)
	r.PUT("/xxxput", req)

	// api 参数
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		// 截取
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})
	// url参数
	r.GET("/userInfo", func(c *gin.Context) {
		name := c.DefaultQuery("name", "默认参数")
		info := c.DefaultQuery("info", "默认参数")
		c.String(http.StatusOK, fmt.Sprintf("hello %s, 默认参数: %s", name, info))
	})
}
