package modules

import (
	"app/dao/mysql"
	"app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type TopNftUser struct {
	UserId     int64     `json:"user_id"`     // 用户id
	Address    string    `json:"address"`     // 用户钱包地址
	Amount     string    `json:"amount"`      // 用户余额
	ReUserId   uint      `json:"re_user_id"`  // 推荐人id
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
}
type AuthInfo struct {
	ID       int64
	Username string
}

func UserLogin(c *gin.Context) {
	// 获取前端传来的数据
	address := c.PostForm("address")
	sign := c.PostForm("sign")
	// 生成签名
	verifySign := tools.Md5(address)
	if sign != verifySign {
		fmt.Println("签名验证出错")
		return
	}
	// 查询数据库
	var user *TopNftUser
	mysql.DB.First(&user).Where("address = ?", address)
	u := AuthInfo{
		ID:       user.UserId,
		Username: "zy",
	}
	token, _ := tools.Sign(tools.AuthInfo(u), "243223ffslsfsldfl412fdsfsdf")
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"info":    user,
			"x-token": token,
		},
	})
}
func GetUserInfo(c *gin.Context) {
	var user *TopNftUser
	mysql.DB.First(&user)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    user,
	})
}
func GetUserList(c *gin.Context) {
	var users *[]TopNftUser
	var count int64
	mysql.DB.Limit(10).Find(&users).Count(&count)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"list":   users,
			"length": len(*users),
			"total":  count,
		},
	})
}
