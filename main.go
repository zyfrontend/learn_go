package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//router.Routers(r)
	r.LoadHTMLGlob("tem/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "message": "谢春阳是傻逼"})
	})
	r.Run(":9797")

	//token := jwt.New(jwt.SigningMethodHS256)
	//token.Claims = jwt.RegisteredClaims{
	//	ID:      "1",
	//	Subject: "zy",
	//}
	//str, err := token.SignedString([]byte("123456"))
	//fmt.Println(str, err)
	//bytes, _ := base64.StdEncoding.DecodeString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
	//fmt.Println(string(bytes))
	//bytes2, _ := base64.StdEncoding.DecodeString("eyJzdWIiOiJ6eSIsImp0aSI6IjEifQ")
	//fmt.Println(string(bytes2))
}
