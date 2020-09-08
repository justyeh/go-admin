package models

import (
	"backend/global"
	"backend/tools"
	"errors"
	"fmt"
)

type Role struct {
	ID       string `json:"id"`
	Name     string `json:"name"  binding:"required"`
	Status   string `json:"status" binding:"required,oneof=active ban"`
	Remark   string `json:"remark"`
	CreateAt int64  `json:"createAt"`
	UpdateAt int64  `json:"updateAt"`
}

type RoleMenu struct {
	RoleId  string   `json:"roleId" binding:"required"`
	MenuIds []string `json:"menuIds" binding:"required"`
}

type RolePermission struct {
	RoleId        string   `json:"roleId" binding:"required"`
	PermissionIds []string `json:"permissionIds" binding:"required"`
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

func (role *Role) RoleMenuIds() ([]string, error) {
	list := []Menu{}
	result := []string{}
	err := global.MYSQL.Raw("SELECT menu_id as id FROM role_menu WHERE role_id = ?", role.ID).Scan(&list).Error
	for _, val := range list {
		result = append(result, val.ID)
	}
	return result, err
}

func (role *Role) RolePermissionIds() ([]string, error) {
	list := []Permission{}
	result := []string{}
	err := global.MYSQL.Raw("SELECT permission_id as id FROM role_permission WHERE role_id = ?", role.ID).Scan(&list).Error
	for _, val := range list {
		result = append(result, val.ID)
	}
	return result, err
}

func (roleMenu *RoleMenu) UpdateRoleMenu() error {
	tx := global.MYSQL.Begin()
	if err := tx.Exec("DELETE FROM role_menu WHERE role_id = ?", roleMenu.RoleId).Error; err != nil {
		tx.Rollback()
		return err
	}
	sql := "INSERT INTO role_menu (id,role_id,menu_id) VALUES "
	for index, menuId := range roleMenu.MenuIds {
		if index == len(roleMenu.MenuIds)-1 {
			sql += fmt.Sprintf("('%s','%s','%s');", tools.UUID(), roleMenu.RoleId, menuId)
		} else {
			sql += fmt.Sprintf("('%s','%s','%s'),", tools.UUID(), roleMenu.RoleId, menuId)
		}
	}
	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (rolePermission *RolePermission) UpdateRolePermission() error {
	tx := global.MYSQL.Begin()
	if err := tx.Exec("DELETE FROM role_permission WHERE role_id = ?", rolePermission.RoleId).Error; err != nil {
		tx.Rollback()
		return err
	}
	sql := "INSERT INTO role_permission (id,role_id,permission_id) VALUES "
	for index, menuId := range rolePermission.PermissionIds {
		if index == len(rolePermission.PermissionIds)-1 {
			sql += fmt.Sprintf("('%s','%s','%s');", tools.UUID(), rolePermission.RoleId, menuId)
		} else {
			sql += fmt.Sprintf("('%s','%s','%s'),", tools.UUID(), rolePermission.RoleId, menuId)
		}
	}
	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
