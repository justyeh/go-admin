package models

import (
	"backend/global"
)

type Menu struct {
	ID       string `json="id"  binding:"required"`
	Name     string `json="name"  binding:"required"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}

func (m *Menu) TableName() string {
	return "menu"
}

func MenuTreeWithName(name string) ([]Menu, error) {
	list := []Menu{}
	err := global.MYSQL.Debug().Where("name like ?", "%"+name+"%").Order("sort asc").Find(&list).Error
	return list, err
}
