package main

import "shopping/common/initialize"

// 测试
func main() {
	initialize.LoadConfig()
	initialize.Mysql()
	initialize.Router()
}
