package initialize

import (
	"backend/global"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitMysql() {
	db, err := gorm.Open("mysql", CONFIG.Mysql.Username+":"+CONFIG.Mysql.Password+"@("+CONFIG.Mysql.Path+")/"+CONFIG.Mysql.Dbname+"?charset=utf8&parseTime=False&loc=Local")
	if err != nil {
		fmt.Println("MySQL启动异常，" + err.Error())
		os.Exit(0)
	}

	db.LogMode(true)
	db.DB().SetMaxIdleConns(CONFIG.Mysql.MaxIdleConns)
	db.DB().SetMaxOpenConns(CONFIG.Mysql.MaxOpenConns)

	global.MYSQL = db
	fmt.Println("数据库连接成功")
}
