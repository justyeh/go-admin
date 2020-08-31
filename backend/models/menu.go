package models

import (
	"backend/global"
)

type Menu struct {
	ID        string `json:"id"  binding:"required"`
	Name      string `json:"name"  binding:"required"`
	Icon      string `json:"icon"`
	Url       string `json:"url"`
	Component string `json:"component"`
	MetaData  string `json:"metaData"`
	Pid       string `json:"pid"`
	Sort      int    `json:"sort"`
	Children  []Menu `json:"children"`
	CreateAt  string `json:"createAt"`
	UpdateAt  string `json:"updateAt"`
}

func (m *Menu) TableName() string {
	return "menu"
}

func (m *Menu) Create() error {
	return global.MYSQL.Create(m).Error
}

func (m *Menu) MenuTreeWithName() ([]Menu, error) {
	list := []Menu{}
	err := global.MYSQL.Where("name like ?", "%"+m.Name+"%").Order("sort asc").Find(&list).Error
	return list, err
}
