package main

import (
	"backend/global"
	"backend/initialize"
)

func main() {
	// 初始化程序功能
	initialize.InitMysql()
	initialize.InitRedis()
	initialize.InitRouter()

	// 程序结束前关闭数据库链接
	defer global.MYSQL.Close()
}
