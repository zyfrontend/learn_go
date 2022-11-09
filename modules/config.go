package modules

import (
	"app/dao/mysql"
	"app/modules/common/response"
	"github.com/gin-gonic/gin"
)

type TopNftConfig struct {
	Group string `json:"group"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// GroupConfig 分组
type GroupConfig struct {
	Publish      map[string]interface{} `json:"publish"`
	Distribution map[string]interface{} `json:"distribution"`
	Withdraw     map[string]interface{} `json:"withdraw"`
	Content      map[string]interface{} `json:"content"`
	Slideshow    map[string]interface{} `json:"slideshow"`
}

func GlobalConfig(c *gin.Context) {
	var config *[]TopNftConfig
	var _ *[]GroupConfig
	mysql.DB.Find(&config)

	// 数据处理
	s := *config
	publish := make(map[string]interface{}, 0)
	withdraw := make(map[string]interface{}, 0)
	distribution := make(map[string]interface{}, 0)
	content := make(map[string]interface{}, 0)
	slideshow := make(map[string]interface{}, 0)
	for a := 0; a < len(s); a++ {
		for b := 1; b < len(s); b++ {
			// 分组
			if s[a].Group == s[b].Group {
				switch s[a].Group {
				case "publish":
					publish[s[a].Key] = s[a].Value
					continue
				case "withdraw":
					withdraw[s[a].Key] = s[a].Value
					continue
				case "distribution":
					distribution[s[a].Key] = s[a].Value
					continue
				case "content":
					content[s[a].Key] = s[a].Value
					continue
				case "slideshow":
					slideshow[s[a].Key] = s[a].Value
					continue
				default:
					break
				}
			}
		}
	}
	response.OkWithData(map[string]interface{}{
		"publish":      publish,
		"slideshow":    slideshow,
		"content":      content,
		"distribution": distribution,
		"withdraw":     withdraw,
		"config":       config,
	}, c)
}
