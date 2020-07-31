package model

import (
	"backend/global"
	"backend/tools"
)

type LoginUser struct {
	Account  string
	Password string
	Captcha  string
	Uuid     string
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
	if err := global.DB.Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}

func (u User) Page(page *tools.Pagination) ([]User, error) {
	list := []User{}
	err := global.DB.Find(&list).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func (u LoginUser) Get() (LoginUser, error) {
	var data LoginUser
	err := global.DB.Find(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

func (u *User) Create() {

}

func (u *User) Update() {

}

func (u *User) Delete() {
	// global.DB.Where("id = ?", u.ID).Delete(User{})
}

func (u *User) IsExist() bool {
	return false
}
