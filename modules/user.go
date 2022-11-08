package modules

import (
	"app/dao/mysql"
	"app/modules/common/response"
	"app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
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
	token := tools.DispatchToken(user.UserId)
	data := map[string]interface{}{
		"info":    user,
		"x-token": token,
	}
	response.OkWithData(data, c)
}
func GetUserInfo(c *gin.Context) {
	address := c.PostForm("address")
	var user *TopNftUser
	mysql.DB.First(&user).Where("address = ?", address)
	response.OkWithData(user, c)
}
func GetUserList(c *gin.Context) {
	var users *[]TopNftUser
	var count int64
	mysql.DB.Limit(10).Find(&users).Count(&count)
	data := map[string]interface{}{
		"list":   users,
		"length": len(*users),
		"total":  count,
	}
	response.OkWithData(data, c)
}
