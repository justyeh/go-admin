package models

import (
	"backend/global"
	"errors"
)

type Role struct {
	ID       string `json:"id"`
	Code     string `json:"code" binding:"required"`
	Name     string `json:"name"  binding:"required"`
	Status   string `json:"status"`
	Pid      string `json:"pid"`
	Sort     int    `json:"sort"`
	Children []Role `json:"children"`
	CreateAt int64  `json:"createAt"`
	UpdateAt int64  `json:"updateAt"`
}

func (role *Role) TableName() string {
	return "role"
}

func (role *Role) RoleList() ([]Role, error) {
	list := []Role{}
	err := global.MYSQL.Where("name LIKE ?", "%"+role.Name+"%").Order("create_at").Find(&list).Error
	return list, err
}

func (role *Role) Create() error {
	var count int
	err := global.MYSQL.Table("role").Where("code = ? ", role.Name).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该角色名称已被使用，无法创建")
	}
	return global.MYSQL.Create(role).Error
}

func (role *Role) Delete() error {
	tx := global.MYSQL.Begin()
	if err := tx.Exec("DELETE FROM user_role WHERE role_id = ?", role.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(role).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (role *Role) UpdateStatus() error {
	return global.MYSQL.Model(role).Update("status", role.Status).Error
}

func (role *Role) Update() error {
	var count int
	err := global.MYSQL.Table("role").Where("id <> ? AND code = ?", role.ID, role.Name).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("修改失败，该角色名称已被占用")
	}
	return global.MYSQL.Model(&role).Updates(map[string]interface{}{
		"code":      role.Code,
		"name":      role.Name,
		"status":    role.Status,
		"pid":       role.Pid,
		"sort":      role.Sort,
		"update_at": role.UpdateAt,
	}).Error
}
