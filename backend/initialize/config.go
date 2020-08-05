package initialize

import (
	"encoding/json"
	"fmt"
	"os"
)

type Mysql struct {
	Username     string
	Password     string
	Path         string
	Dbname       string
	MaxIdleConns int
	MaxOpenConns int
}

type Redis struct {
	Username string
	Password string
}

type Gin struct {
	Mode string
	Port string
}

type Config struct {
	Mysql Mysql
	Redis Redis
	Gin   Gin
}

var (
	CONFIG Config
)

func init() {
	file, err := os.Open("./config.json")
	if err != nil {
		panic("配置文件不存在")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if decoder.Decode(&CONFIG) != nil {
		panic("读取配置文件失败")
	}

	fmt.Println("配置文件解析完成")
}
