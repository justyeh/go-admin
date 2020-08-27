package models

type Dept struct {
	ID       string `json="id"  binding:"required"`
	Name     string `json="name"  binding:"required"`
	Status   string `json="sort"  binding:"required"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}
