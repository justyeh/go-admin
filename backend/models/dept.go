package models

import (
	"backend/global"
	"errors"
)

type Dept struct {
	ID       string `json:"id"`
	Name     string `json:"name"  binding:"required"`
	Pid      string `json:"pid"`
	Sort     int    `json:"sort"`
	Children []Dept `json:"children"`
	CreateAt int64  `json:"createAt"`
	UpdateAt int64  `json:"updateAt"`
}

func (dept *Dept) TableName() string {
	return "dept"
}

func (dept *Dept) DeptTree() ([]Dept, error) {
	list := []Dept{}
	var err error
	if len(dept.Name) == 0 {
		err = global.MYSQL.Order("sort DESC").Order("create_at").Find(&list).Error
	} else {
		err = global.MYSQL.Where("name LIKE ?", "%"+dept.Name+"%").Order("sort DESC").Order("create_at").Find(&list).Error
	}
	return list, err
}

func (dept *Dept) Create() error {
	var count int
	err := global.MYSQL.Table("dept").Where("name = ? ", dept.Name).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该部门名称已被使用，无法创建")
	}
	return global.MYSQL.Create(dept).Error
}

func (dept *Dept) Delete() error {
	var count int
	err := global.MYSQL.Table("user").Where("dept_id = ? ", dept.ID).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("删除失败，该部门存在关联用户")
	}
	return global.MYSQL.Delete(dept).Error
}

func (dept *Dept) Update() error {
	var count int
	err := global.MYSQL.Table("dept").Where("id <> ? AND name = ?", dept.ID, dept.Name).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("修改失败，该部门名称已被占用")
	}
	return global.MYSQL.Model(&dept).Updates(map[string]interface{}{
		"name":      dept.Name,
		"pid":       dept.Pid,
		"sort":      dept.Sort,
		"update_at": dept.UpdateAt,
	}).Error
}
