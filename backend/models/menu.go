package models

import (
	"backend/global"
	"errors"
)

type Menu struct {
	ID        string `json:"id"`
	Name      string `json:"name"  binding:"required"`
	Icon      string `json:"icon"`
	Url       string `json:"url"`
	Component string `json:"component"`
	MetaData  string `json:"metaData"`
	Pid       string `json:"pid"`
	Sort      int    `json:"sort"`
	Children  []Menu `json:"children"`
	CreateAt  int64  `json:"createAt"`
	UpdateAt  int64  `json:"updateAt"`
}

func (menu *Menu) TableName() string {
	return "menu"
}

func (menu *Menu) MenuById() error {
	return global.MYSQL.Where("id = ?", menu.ID).First(&menu).Error
}

func (menu *Menu) MenuTreeWithName() ([]Menu, error) {
	list := []Menu{}
	err := global.MYSQL.Where("name like ?", "%"+menu.Name+"%").Order("sort asc").Order("create_at").Find(&list).Error
	return list, err
}

func (menu *Menu) Create() error {
	var count int
	err := global.MYSQL.Table("menu").Where("name = ? ", menu.Name).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该菜单名称已被使用，无法创建")
	}
	return global.MYSQL.Create(menu).Error
}

func (menu *Menu) Delete() error {
	var count int
	err := global.MYSQL.Table("role_menu").Where("menu_id = ? ", menu.ID).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("删除失败，该菜单存在关联角色")
	}
	return global.MYSQL.Delete(menu).Error
}

func (menu *Menu) Update() error {
	var count int
	err := global.MYSQL.Table("menu").Where("id != ? and name = ?", menu.ID, menu.Name).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("创建失败，该菜单名称已被占用")
	}
	return global.MYSQL.Model(&menu).Updates(map[string]interface{}{
		"name":      menu.Name,
		"icon":      menu.Icon,
		"url":       menu.Url,
		"component": menu.Component,
		"metaData":  menu.Component,
		"pid":       menu.Pid,
		"sort":      menu.Sort,
		"update_at": menu.UpdateAt,
	}).Error
}
