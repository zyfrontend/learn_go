package main

import (
	"fmt"
)

func main() {
	//r := gin.Default()
	//router.Routers(r)
	//
	//r.Run(":9797")
	token := DispatchToken(0226)
	fmt.Printf("当前生成token为%s\n", token)
	VerifyToken(token)
}
