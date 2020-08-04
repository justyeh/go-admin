package main

import (
	"backend/global"
	"backend/initialize"
	"fmt"
)

func main() {
	// 程序初始化
	initialize.InitMysql()
	initialize.InitRedis()
	initialize.InitRouter()

	// 程序结束前关闭数据库链接
	defer global.MYSQL.Close()

	fmt.Println("程序启动成功，服务运行与http://127.0.0.1:1234")
}
