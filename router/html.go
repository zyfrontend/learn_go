package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestHtml(r *gin.Engine) {
	r.LoadHTMLGlob("tem/*")

	r.GET("/index", func(c *gin.Context) {
		// c.HTML(http.StatusOK, "index.html", gin.H{"title": "测试title", "ce": "1111"})
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "address": "www.5lmh.com"})

	})
}
