package tools

func IsExistInSlice(source []interface{}, target interface{}) bool {
	for _, val := range source {
		if target == val {
			return true
		}
	}
	return false
}
