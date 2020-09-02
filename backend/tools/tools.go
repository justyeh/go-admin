package tools

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// 判断切片中是否存在某个元素
func IsExistInSlice(source []interface{}, target interface{}) bool {
	for _, val := range source {
		if target == val {
			return true
		}
	}
	return false
}

// 生成UUID
func UUID() string {
	return uuid.NewV1().String()
}

// 获取当前时间戳
func GetUnixNow() int64 {
	return time.Now().Unix()
}
