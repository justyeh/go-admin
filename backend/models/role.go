package models

import (
	"backend/global"
	"backend/tools"
	"errors"
)

type Role struct {
	ID       string `json:"id"`
	Name     string `json:"name"  binding:"required"`
	Status   string `json:"status" binding:"required,oneof=active ban"`
	Remark   string `json:"remark"`
	CreateAt int64  `json:"createAt"`
	UpdateAt int64  `json:"updateAt"`
}

func (role *Role) TableName() string {
	return "role"
}

func (role *Role) RoleList(page tools.Pagination) ([]Role, int, error) {
	var list = []Role{}
	var total int
	var err error

	db := global.MYSQL
	if len(role.Name) > 0 {
		db = db.Where("name LIKE ?", "%"+role.Name+"%")
	}
	err = db.Table("role").Count(&total).Error
	if err != nil || total == 0 {
		return list, total, err
	}
	err = db.Order("create_at DESC").Offset((page.Current - 1) * page.Size).Limit(page.Size).Find(&list).Error
	return list, total, err
}

func (role *Role) Create() error {
	var count int
	err := global.MYSQL.Table("role").Where("name = ? ", role.Name).Count(&count).Error
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
	if err := tx.Exec("DELETE FROM role_menu WHERE role_id = ?", role.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Exec("DELETE FROM role_permission WHERE role_id = ?", role.ID).Error; err != nil {
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
	return global.MYSQL.Model(&role).Updates(map[string]interface{}{
		"status":    role.Status,
		"update_at": role.UpdateAt,
	}).Error
}

func (role *Role) Update() error {
	var count int
	err := global.MYSQL.Table("role").Where("id <> ? AND name = ?", role.ID, role.Name).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("修改失败，该角色名称已被占用")
	}
	return global.MYSQL.Model(&role).Updates(map[string]interface{}{
		"name":      role.Name,
		"status":    role.Status,
		"remark":    role.Remark,
		"update_at": role.UpdateAt,
	}).Error
}

func (role *Role) RoleMenuList() ([]Menu, error) {
	list := []Menu{}
	err := global.MYSQL.Raw("SELECT menu.id FROM menu,role_menu WHERE role_menu.role_id = ?", role.ID).Scan(&list).Error
	return list, err
}

func (role *Role) RolePermisssionList() ([]Permission, error) {
	list := []Permission{}
	err := global.MYSQL.Raw("SELECT permission.id FROM permission,role_permission WHERE role_permission.role_id = ?", role.ID).Scan(&list).Error
	return list, err
}

func (role *Role) UpdateRoleMenu() ([]Menu, error) {
	list := []Menu{}
	err := global.MYSQL.Raw("SELECT menu.id FROM menu,role_menu WHERE role_menu.role_id = ?", role.ID).Scan(&list).Error
	return list, err
}

func (role *Role) UpdateRolePermisssion() ([]Permission, error) {
	list := []Permission{}
	err := global.MYSQL.Raw("SELECT permission.id FROM permission,role_permission WHERE role_permission.role_id = ?", role.ID).Scan(&list).Error
	return list, err
}
