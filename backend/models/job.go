package models

type Job struct {
	ID       string `json="id"  binding:"required"`
	Name     string `json="name"  binding:"required"`
	Sort     string `json="sort"  binding:"required"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}
