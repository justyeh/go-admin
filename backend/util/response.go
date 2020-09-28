package util

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(c *gin.Context, data gin.H) {
	data["status"] = http.StatusOK
	c.JSON(http.StatusOK, data)
}

func ResponseError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusBadRequest,
		"message": message,
	})
}

func getTranslateError(s string) string {
	arr := strings.Split(s, "'")

	if len(arr) < 6 {
		return s
	}

	field := arr[3]
	tag := arr[5]

	switch tag {
	case "required":
		return field + "字段为必填项"
	case "email":
		return field + "字段邮箱格式错误"
	case "oneof":
		return field + "字段的值不符合要求"
	default:
		return s
	}
}

func ResponseBindError(c *gin.Context, err error) {
	message := []string{}

	errs := strings.Split(err.Error(), "\n")
	for _, val := range errs {
		message = append(message, getTranslateError(val))
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusBadRequest,
		"message": strings.Join(message, "; "),
	})
}
