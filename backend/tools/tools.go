package tools

// 判断切片中是否存在某个元素
func IsExistInSlice(source []interface{}, target interface{}) bool {
	for _, val := range source {
		if target == val {
			return true
		}
	}
	return false
}

// 切片转为树形结构
func SliceToTree(source []interface{}) []interface {
}
