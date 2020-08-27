package models

import (
	"backend/global"
	"backend/tools"
)

type Login struct {
	Account  string `json:"account"  binding:"required,email"`
	Password string `json:"password"  binding:"required"`
	Captcha  string `json:"captcha"  binding:"required"`
	Uuid     string `json:"uuid"  binding:"required"`
}

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
	ID  string `json:"id"`
	Url string `json:"url"`
}

type User struct {
	ID         string           `json:"id"`
	Account    string           `json:"account"`
	Password   string           `json:"password"`
	Status     string           `json:"status"`
	Nickname   string           `json:"nickname"`
	Phone      string           `json:"phone"`
	Email      string           `json:"email"`
	Dept       UserDept         `json:"dept"`
	Job        UserJob          `json:"job"`
	Menu       []UserMenu       `json:"menu"`
	Permission []UserPermission `json:"permission" `
	CreateAt   int              `json:"createAt"`
	UpdateAt   int              `json:"updateAt"`
}

func (u LoginUser) TableName() string {
	return "user"
}

func (u User) TableName() string {
	return "user"
}

func (u *LoginUser) Login() {
	db := global.MYSQL.Where("account = ? and password = ?", u.Account, u.Password).First(&u)
}

func (u *User) UserInfo() {
	db := global.MYSQL.Where("id = ?", u.ID).First(&u)
	permissionList := []UserPermission{}
	db.Related(&permissionList)

	menuList := []UserMenu{}
	db.Related(&menuList)
}

func (u User) List() ([]User, error) {
	list := []User{}
	if err := global.MYSQL.Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}

func (u User) Page(page *tools.Pagination) ([]User, error) {
	list := []User{}
	err := global.MYSQL.Find(&list).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func (u *User) Create() (string, error) {
	return "", nil
}

func (u *User) Update() {

}

func (u *User) Delete() {
	// global.MYSQL.Where("id = ?", u.ID).Delete(User{})
}

func (u *User) IsExist() bool {
	return false
}
