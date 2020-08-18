package models

import (
	"backend/global"
	"backend/tools"
)

type Login struct {
	Account  string `json:"account"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Captcha  string `json:"captcha"  binding:"required"`
	Uuid     string `json:"uuid"  binding:"required"`
}

type User struct {
	ID   int64  `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}

func (u User) TableName() string {
	return "user"
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

func (u Login) Get() (Login, error) {
	var data Login
	err := global.MYSQL.Find(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
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
