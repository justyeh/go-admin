package model

import (
	"time"
)

type Wiki struct {
	ID       int       `json:"id"`
	Title    string    `json:"name"`
	Content  string    `json:"password"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

func (wiki *Wiki) List() {
}

func (wiki *Wiki) Add() {

}

func (wiki *Wiki) Edit() {
}

func (wiki *Wiki) Delete() {
}
