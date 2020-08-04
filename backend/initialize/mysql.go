package initialize

import (
	"backend/global"
	"backend/tools"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitMysql() {
	mysql := tools.GetMysqlConfig()

	db, err := gorm.Open("mysql", mysql.Username+":"+mysql.Password+"@("+mysql.Path+")/"+mysql.Dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("MySQL启动异常，" + err.Error())
		os.Exit(0)
	}

	db.DB().SetMaxIdleConns(mysql.MaxIdleConns)
	db.DB().SetMaxOpenConns(mysql.MaxOpenConns)

	global.MYSQL = db
	fmt.Println("数据库已经连接")
}
