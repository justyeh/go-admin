package models

import (
	"backend/global"
	"backend/tools"
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
	Account    string           `json:"account"`
	Status     string           `json:"status"`
	Nickname   string           `json:"nickname"`
	Phone      string           `json:"phone"`
	Email      string           `json:"email"`
	Dept       UserDept         `json:"dept"`
	Job        UserJob          `json:"job"`
	Menu       []UserMenu       `json:"menu"`
	Permission []UserPermission `json:"permission" `
	CreateAt   tools.TimeStamp  `json:"createAt"`
	UpdateAt   tools.TimeStamp  `json:"updateAt"`
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
