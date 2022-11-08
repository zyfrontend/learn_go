package tools

import (
	"github.com/spf13/viper"
)

var Search = InitConfig

func InitConfig(path string) string {
	//path, err := os.Getwd()
	//if err != nil {
	//	panic(err)
	//}
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	p := viper.GetString(path)
	return p
}
