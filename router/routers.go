package router

import (
	"app/middleware"
	"app/modules"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	// 非鉴权接口
	v1 := r.Group("/api")
	{
		//登录
		v1.POST("/login", modules.UserLogin)
		v1.GET("/config", modules.GlobalConfig)
	}

	// 鉴权接口
	v2 := r.Group("/auth")
	v2.Use(middleware.VerifyPermissions())
	{
		// 用户信息
		v2.GET("/userInfo", modules.GetUserInfo)
		v2.GET("/userList", modules.GetUserList)
		v2.GET("/pushList", modules.GetPushList)
		v2.GET("/orderList", modules.GetOrderList)
	}
}
