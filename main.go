package main

import (
	"App/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 注册使用中间件

	// 全局中间件
	// r.Use(middleware.MiddleWare())
	// 路由分包到单文件
	// r := router.Upload()

	// 路由分包到多文件
	router.Get(r)
	router.Uploads(r)
	router.LoginJSON(r)
	router.LoginFORM(r)
	router.TestResponse(r)
	router.TestHtml(r)
	router.TestMiddleWare(r)
	router.Cookie(r)
	if err := r.Run(":8000"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
