package main

import (
	"app/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Routers(r)

	r.Run(":9797")
	//token := tools.DispatchToken(0226)
	//fmt.Printf("当前生成token为%s\n", token)
	//tools.VerifyToken(token)
}
