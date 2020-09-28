package util

import (
	"backend/global"
	"crypto/md5"
	"encoding/hex"
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
func IsExistInStringSlice(source []string, target interface{}) bool {
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

// MD5
func GetMD5(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

type tree struct {
	ID string
}

func GetChildIds(tableName string, pid string) []string {
	childList := []tree{}
	recursionSelect(tableName, []tree{{ID: pid}}, &childList)
	result := []string{}
	for _, val := range childList {
		result = append(result, val.ID)
	}
	return result
}

func recursionSelect(tableName string, pIds []tree, childList *[]tree) {
	for _, val := range pIds {
		ids := []tree{}
		if err := global.MYSQL.Table(tableName).Where("pid = ?", val.ID).Scan(&ids).Error; err != nil {
			break
		} else {
			*childList = append(*childList, ids...)
			recursionSelect(tableName, ids, childList)
		}
	}
}
