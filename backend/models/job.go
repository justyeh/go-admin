package models

import (
	"backend/global"
	"backend/util"
	"errors"
)

type Job struct {
	ID       string `json:"id"`
	Name     string `json:"name"  binding:"required"`
	Sort     int    `json:"sort"`
	CreateAt int64  `json:"createAt"`
	UpdateAt int64  `json:"updateAt"`
}

func (job *Job) TableName() string {
	return "job"
}

func (job *Job) JobList(page util.Pagination) ([]Job, int, error) {
	var list = []Job{}
	var total int
	var err error

	db := global.MYSQL.Table("job")
	if len(job.Name) > 0 {
		db = db.Where("name LIKE ?", "%"+job.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil || total == 0 {
		return list, total, err
	}
	err = db.Order("sort ASC").Order("create_at DESC").Offset((page.Current - 1) * page.Size).Limit(page.Size).Find(&list).Error
	return list, total, err
}

func (job *Job) Create() error {
	var count int
	err := global.MYSQL.Table("job").Where("name = ? ", job.Name).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该岗位名称已被使用，无法创建")
	}
	return global.MYSQL.Create(job).Error
}

func (job *Job) Delete() error {
	var count int
	err := global.MYSQL.Table("user").Where("job_id = ? ", job.ID).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("删除失败，该岗位存在关联用户")
	}
	return global.MYSQL.Delete(job).Error
}

func (job *Job) Update() error {
	var count int
	err := global.MYSQL.Table("job").Where("id <> ? AND name = ?", job.ID, job.Name).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("修改失败，该岗位名称已被占用")
	}
	return global.MYSQL.Omit("create_at").Save(job).Error
}
