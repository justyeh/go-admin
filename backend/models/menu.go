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

func (menu *Menu) MenuTree() ([]Menu, error) {
	list := []Menu{}
	db := global.MYSQL
	if len(menu.Name) > 0 {
		db = db.Where("name LIKE ?", "%"+menu.Name+"%")
	}
	err := db.Order("sort ASC").Order("create_at DESC").Find(&list).Error
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

func (menu *Menu) Delete(ids []string) error {
	tx := global.MYSQL.Begin()
	if err := tx.Exec("DELETE FROM role_menu WHERE menu_id IN (?)", ids).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Exec("DELETE FROM menu WHERE id IN (?)", ids).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (menu *Menu) Update() error {
	var count int
	err := global.MYSQL.Table("menu").Where("id <> ? AND name = ?", menu.ID, menu.Name).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("修改失败，该菜单名称已被占用")
	}
	return global.MYSQL.Omit("create_at").Save(menu).Error
}
