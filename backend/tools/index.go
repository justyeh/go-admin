package tools

import (
	"strconv"
)

//StringToInt 字符串转数字int
func StringToInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		HasError(err, "string to int error"+err.Error(), -1)
	}
	return result
}

// HasError 错误断言
// 当 error 不为 nil 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
// 若 msg 为空，则默认为 error 中的内容
func HasError(err error, msg string, code ...int) {
	if err != nil {
		statusCode := 200
		if len(code) > 0 {
			statusCode = code[0]
		}
		if msg == "" {
			msg = err.Error()
		}
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}
