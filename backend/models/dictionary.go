package models

import (
	"backend/global"
	"backend/tools"
	"errors"
)

type Dictionary struct {
	ID          string `json:"id"`
	Name        string `json:"name"  binding:"required"`
	Description string `json:"description"`
	CreateAt    int64  `json:"createAt"`
	UpdateAt    int64  `json:"updateAt"`
}

type DictionaryDetail struct {
	ID           string `json:"id"`
	DictionaryId string `json:"dictionaryId"  binding:"required"`
	Label        string `json:"label"  binding:"required"`
	Value        string `json:"value"  binding:"required"`
	Sort         int    `json:"sort"`
	CreateAt     int64  `json:"createAt"`
	UpdateAt     int64  `json:"updateAt"`
}

func (dictionary *Dictionary) TableName() string {
	return "dictionary"
}

func (dictionaryDetail *DictionaryDetail) TableName() string {
	return "dictionary_detail"
}

func (dictionary *Dictionary) DictionaryList(page tools.Pagination) ([]Dictionary, int, error) {
	var list = []Dictionary{}
	var total int
	var err error

	db := global.MYSQL.Table("dictionary")
	if len(dictionary.Name) > 0 {
		db = db.Where("name LIKE ?", "%"+dictionary.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil || total == 0 {
		return list, total, err
	}
	err = db.Order("create_at DESC").Offset((page.Current - 1) * page.Size).Limit(page.Size).Find(&list).Error
	return list, total, err
}

func (dictionary *Dictionary) Create() error {
	var count int
	err := global.MYSQL.Table("dictionary").Where("name = ? ", dictionary.Name).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该字典名称已被使用，无法创建")
	}
	return global.MYSQL.Create(dictionary).Error
}

func (dictionary *Dictionary) Delete() error {
	tx := global.MYSQL.Begin()
	if err := tx.Exec("DELETE FROM dictionary_detail WHERE dictionary_id = ?", dictionary.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := global.MYSQL.Delete(dictionary).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (dictionary *Dictionary) Update() error {
	var count int
	err := global.MYSQL.Table("dictionary").Where("id <> ? AND name = ?", dictionary.ID, dictionary.Name).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("修改失败，该字典名称已被占用")
	}
	return global.MYSQL.Omit("create_at").Save(dictionary).Error
}

func (dictionaryDetail *DictionaryDetail) DictionaryDetailList() ([]DictionaryDetail, error) {
	var list = []DictionaryDetail{}
	err := global.MYSQL.Where("dictionary_id = ?", dictionaryDetail.DictionaryId).Order("sort ASC").Order("create_at DESC").Find(&list).Error
	return list, err
}

func (dictionaryDetail *DictionaryDetail) Create() error {
	var count int
	parmas := []interface{}{dictionaryDetail.DictionaryId, dictionaryDetail.Label, dictionaryDetail.Value}
	err := global.MYSQL.Table("dictionary_detail").Where("dictionary_id = ? And (label = ? or value = ?)", parmas...).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该字典标签/值已被使用，无法创建")
	}
	return global.MYSQL.Create(dictionaryDetail).Error
}

func (dictionaryDetail *DictionaryDetail) Delete() error {
	return global.MYSQL.Delete(dictionaryDetail).Error
}

func (dictionaryDetail *DictionaryDetail) Update() error {
	var count int
	parmas := []interface{}{dictionaryDetail.DictionaryId, dictionaryDetail.ID, dictionaryDetail.Label, dictionaryDetail.Value}
	err := global.MYSQL.Table("dictionary_detail").Where("dictionary_id = ? AND id <> ? AND (label = ? OR value = ?)", parmas...).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("修改失败，该字典标签/值已被占用")
	}
	return global.MYSQL.Omit("create_at").Save(dictionaryDetail).Error
}
