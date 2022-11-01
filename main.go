package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
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

	// 表单参数
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("userName")
		password := c.PostForm("userPassword")
		c.String(http.StatusOK, fmt.Sprintf("username: %s, password: %s, type: %s", username, password, types))
	})

	// 单文件上传
	// 限制上传大小
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "图片上传出错")
			panic(err)
		}
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})

	// 上传特定文件
	r.POST("/uploadFile", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file: %v", err)
		}
		// 获取文件大小
		if headers.Size > 1024*1024*2 {
			fmt.Println("文件太大了")
			return
		}
		// 上传文件的类型
		if headers.Header.Get("Content-Type") != "image/png" {
			fmt.Println("只允许上传png图片")
			return
		}
		c.SaveUploadedFile(headers, "./img/"+headers.Filename)
		c.String(http.StatusOK, headers.Filename)
	})

	// 多文件上传
	r.POST("/uploadFiles", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有图片
		files := form.File["files"]
		// 遍历所有图片
		for _, file := range files {
			// 逐个存
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})

	r.Run(":8000")
}
