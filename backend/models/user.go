package models

import (
	"backend/global"
	"backend/util"
	"errors"
	"fmt"
)

type LoginUser struct {
	ID       string `json:"id"`
	Account  string `json:"account"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Captcha  string `json:"captcha"  binding:"required"`
	Uuid     string `json:"uuid"  binding:"required"`
}

type User struct {
	ID       string `json:"id"`
	Account  string `json:"account" binding:"required"`
	Password string `json:"-"`
	Nickname string `json:"nickname" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Status   string `json:"status" binding:"required,oneof=active ban"`
	CreateAt int64  `json:"createAt"`
	UpdateAt int64  `json:"updateAt"`
}

type ViewUser struct {
	User
	DeptId   string `json:"deptId"`
	JobId    string `json:"jobId"`
	DeptName string `json:"deptName"`
	JobName  string `json:"jobName"`
}

type SystemUser struct {
	User
	DeptId  string   `json:"deptId" binding:"required"`
	JobId   string   `json:"jobId" binding:"required"`
	RoleIds []string `json:"roleIds" binding:"required" sql:"-"`
}

type UserStatus struct {
	ID       string `json:"id"  binding:"required"`
	Status   string `json:"status" binding:"required,oneof=active ban"`
	UpdateAt int64  `json:"updateAt"`
}

func (l *LoginUser) UserWithAccountAndPassword() error {
	err := global.MYSQL.Table("user").Where("account = ? AND password = ?", l.Account, l.Password).First(&l).Error
	return err
}

func (u *User) UserInfoWithID() error {
	/* db := global.MYSQL.Where("id = ?", u.ID).First(&u)
	prmissionList := []UserPermission{}
	db.Related(&prmissionList)
	return db.Error */
	return nil
}

func (user *ViewUser) UserList(page util.Pagination) ([]ViewUser, int, error) {
	list := []ViewUser{}
	var total int
	var err error

	db := global.MYSQL.Table("user")
	db = db.Select("user.*, dept.name as dept_name, job.name as job_name")
	db = db.Joins("LEFT JOIN dept ON user.dept_id = dept.id")
	db = db.Joins("LEFT JOIN job ON user.job_id = job.id")
	db = db.Where("user.id <> '0'")
	if len(user.Account) > 0 || len(user.Nickname) > 0 {
		db = db.Where("(user.account LIKE ? OR user.nickname LIKE ?)", "%"+user.Account+"%", "%"+user.Nickname+"%")
	}
	if len(user.DeptId) > 0 {
		db = db.Where("user.dept_id = ?", user.DeptId)
	}
	err = db.Count(&total).Error
	if err != nil || total == 0 {
		return list, total, err
	}
	err = db.Order("user.create_at DESC").Offset((page.Current - 1) * page.Size).Limit(page.Size).Find(&list).Error
	return list, total, err
}

func (user *SystemUser) Create() error {
	var count int
	err := global.MYSQL.Table("user").Where("account = ? ", user.Account).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该账号已被使用，无法创建")
	}

	tx := global.MYSQL.Begin()
	// user表插入记录
	if err := tx.Table("user").Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	// user_role表先清除相关记录
	if err := tx.Exec("DELETE FROM user_role WHERE user_id = ?", user.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	// user_role插入记录
	sql := "INSERT INTO user_role (id,user_id,role_id) VALUES "
	for index, roleId := range user.RoleIds {
		if index == len(user.RoleIds)-1 {
			sql += fmt.Sprintf("('%s','%s','%s');", util.UUID(), user.ID, roleId)
		} else {
			sql += fmt.Sprintf("('%s','%s','%s'),", util.UUID(), user.ID, roleId)
		}
	}
	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (user *User) Delete() error {
	tx := global.MYSQL.Begin()
	if err := tx.Exec("DELETE FROM user_role WHERE user_id = ?", user.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Table("user").Delete(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (us *UserStatus) UpdateStatus() error {
	return global.MYSQL.Table("user").Save(us).Error
}

func (user *SystemUser) Update() error {
	var count int
	err := global.MYSQL.Table("user").Where("id <> ? AND account = ?", user.ID, user.Account).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("修改失败，该账号名称已被占用")
	}

	tx := global.MYSQL.Begin()
	// user表更新记录

	if err = tx.Table("user").Omit("create_at").Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// user_role表先清除相关记录
	if err := tx.Exec("DELETE FROM user_role WHERE user_id = ?", user.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	// user_role插入记录
	sql := "INSERT INTO user_role (id,user_id,role_id) VALUES "
	for index, roleId := range user.RoleIds {
		if index == len(user.RoleIds)-1 {
			sql += fmt.Sprintf("('%s','%s','%s');", util.UUID(), user.ID, roleId)
		} else {
			sql += fmt.Sprintf("('%s','%s','%s'),", util.UUID(), user.ID, roleId)
		}
	}
	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
