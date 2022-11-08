package modules

import (
	"app/dao/mysql"
	"app/modules/common/response"
	"github.com/gin-gonic/gin"
)

type TopNftConfig struct {
	Group string `json:"group"`
	Value string `json:"value"`
}

func GlobalConfig(c *gin.Context) {
	var config *[]TopNftConfig
	mysql.DB.Find(&config)
	response.OkWithData(config, c)
}
