package models

import (
	"backend/global"
	"backend/tools"
	"errors"
)

type UserDept struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserJob struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserPermission struct {
	Code string `json:"code"`
}

type UserMenu struct {
	ID string
}

type LoginUser struct {
	ID       string `json:"id"`
	Account  string `json:"account"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Captcha  string `json:"captcha"  binding:"required"`
	Uuid     string `json:"uuid"  binding:"required"`
}

type User struct {
	ID         string           `json:"id"`
	Account    string           `json:"account" binding:"required"`
	Password   string           `json:"-"`
	Nickname   string           `json:"nickname" binding:"required"`
	Phone      string           `json:"phone" binding:"required"`
	Email      string           `json:"email" binding:"required"`
	Status     string           `json:"status" binding:"required"`
	Dept       UserDept         `json:"dept"`
	Job        UserJob          `json:"job"`
	Menu       []UserMenu       `json:"menu"`
	Permission []UserPermission `json:"permission"`
	CreateAt   int64            `json:"createAt"`
	UpdateAt   int64            `json:"updateAt"`
}

func (l LoginUser) TableName() string {
	return "user"
}

func (u User) TableName() string {
	return "user"
}

func (l *LoginUser) UserWithAccountAndPassword() error {
	err := global.MYSQL.Where("account = ? AND password = ?", l.Account, l.Password).First(&l).Error
	return err
}

func (u *User) UserInfoWithID() error {
	db := global.MYSQL.Where("id = ?", u.ID).First(&u)
	prmissionList := []UserPermission{}
	db.Related(&prmissionList)
	return db.Error
}

func (user *User) UserList(page tools.Pagination) ([]User, int, error) {
	list := []User{}
	var total int
	var err error

	db := global.MYSQL
	if len(user.Account) > 0 || len(user.Nickname) > 0 {
		db = db.Where("id <> '0' AND (account LIKE ? OR nickname LIKE ?)", "%"+user.Account+"%", "%"+user.Nickname+"%")
	} else {
		db = db.Where("id <> '0'")
	}

	err = db.Table("user").Count(&total).Error
	if err != nil || total == 0 {
		return list, total, err
	}

	err = db.Order("create_at DESC").Offset((page.Current - 1) * page.Size).Limit(page.Size).Find(&list).Error
	return list, total, err
}

func (user *User) Create() error {
	var count int
	err := global.MYSQL.Table("user").Where("account = ? ", user.Account).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该账号已被使用，无法创建")
	}
	return global.MYSQL.Create(user).Error
}

func (user *User) Delete() error {
	tx := global.MYSQL.Begin()
	if err := tx.Exec("DELETE FROM user_user WHERE user_id = ?", user.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (user *User) UpdateStatus() error {
	return global.MYSQL.Model(&user).Updates(map[string]interface{}{
		"status":    user.Status,
		"update_at": user.UpdateAt,
	}).Error
}

func (user *User) Update() error {
	var count int
	err := global.MYSQL.Table("user").Where("id <> ? AND account = ?", user.ID, user.Account).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("修改失败，该账号名称已被占用")
	}
	return global.MYSQL.Model(&user).Updates(map[string]interface{}{
		"account":   user.Account,
		"nickname":  user.Nickname,
		"email":     user.Nickname,
		"phone":     user.Nickname,
		"status":    user.Status,
		"update_at": user.UpdateAt,
	}).Error
}
