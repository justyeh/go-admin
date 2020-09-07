package tools

import (
	"backend/global"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
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

// MD5
func GetMD5(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// 同样的body数据做缓存处理
var BodyData string
var BodyMap map[string]string

// 获取body数据
func GetBodyData(c *gin.Context, key string) string {
	data, err := c.GetRawData()
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // important：把读过的字节流重新放到body
	if err != nil {
		return ""
	}

	m := make(map[string]string)
	if string(data) == BodyData {
		m = BodyMap
	} else {
		if err := json.Unmarshal(data, &m); err != nil {
			return ""
		}
		BodyData = string(data)
		BodyMap = m
	}

	for k, v := range m {
		if k == key {
			return v
		}
	}
	return ""

}

func GetChildIds(tableName string, pid string) []string {
	result := []string{}
	recursionSelect(tableName, []string{pid}, &result)
	return result
}

func recursionSelect(tableName string, pIds []string, result *[]string) {
	for _, val := range pIds {
		ids := []string{}
		if err := global.MYSQL.Exec("SELECT id FROM ? WHERE pid = ", tableName, val).Scan(&ids).Error; err != nil {
			break
		} else {
			*result = append(*result, ids...)
			recursionSelect(tableName, ids, result)
		}
	}
}
