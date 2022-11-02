package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Uploads(r *gin.Engine) {
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
}

// func Upload() *gin.Engine {
// 	r := gin.Default()
// 	// 单文件上传
// 	// 限制上传大小
// 	r.MaxMultipartMemory = 8 << 20
// 	r.POST("/upload", func(c *gin.Context) {
// 		file, err := c.FormFile("file")
// 		if err != nil {
// 			c.String(500, "图片上传出错")
// 			panic(err)
// 		}
// 		c.SaveUploadedFile(file, file.Filename)
// 		c.String(http.StatusOK, file.Filename)
// 	})

// 	// 上传特定文件
// 	r.POST("/uploadFile", func(c *gin.Context) {
// 		_, headers, err := c.Request.FormFile("file")
// 		if err != nil {
// 			log.Printf("Error when try to get file: %v", err)
// 		}
// 		// 获取文件大小
// 		if headers.Size > 1024*1024*2 {
// 			fmt.Println("文件太大了")
// 			return
// 		}
// 		// 上传文件的类型
// 		if headers.Header.Get("Content-Type") != "image/png" {
// 			fmt.Println("只允许上传png图片")
// 			return
// 		}
// 		c.SaveUploadedFile(headers, "./img/"+headers.Filename)
// 		c.String(http.StatusOK, headers.Filename)
// 	})

// 	// 多文件上传
// 	r.POST("/uploadFiles", func(c *gin.Context) {
// 		form, err := c.MultipartForm()
// 		if err != nil {
// 			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
// 		}
// 		// 获取所有图片
// 		files := form.File["files"]
// 		// 遍历所有图片
// 		for _, file := range files {
// 			// 逐个存
// 			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
// 				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
// 				return
// 			}
// 		}
// 		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
// 	})
// 	return r
// }
