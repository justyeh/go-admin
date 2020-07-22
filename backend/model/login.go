package model

//Login 结构体
type Login struct {
	Name string `json:"name"`
	Password string  `json:"password"`
	Code string  `json:"code"`
	UUID string `json:"uuid"`
}
