package app

import (
	"backend/models"
	"backend/util"

	"github.com/gin-gonic/gin"
)

func JobList(c *gin.Context) {
	job := models.Job{Name: c.Query("keyword")}
	page := util.NewPagination(c)

	list, total, err := job.JobList(page)
	if err != nil {
		util.ResponseError(c, err.Error())
		return
	}

	util.ResponseSuccess(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func AddJob(c *gin.Context) {
	now := util.GetUnixNow()
	job := models.Job{CreateAt: now, UpdateAt: now}
	if err := c.ShouldBind(&job); err != nil {
		util.ResponseBindError(c, err)
		return
	}

	job.ID = util.UUID()

	if err := job.Create(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "添加成功", "data": job})
}

func EditJob(c *gin.Context) {
	job := models.Job{UpdateAt: util.GetUnixNow()}
	if err := c.ShouldBind(&job); err != nil {
		util.ResponseBindError(c, err)
		return
	}

	if err := job.Update(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func DeleteJob(c *gin.Context) {
	job := models.Job{ID: c.Param("id")}

	if len(job.ID) == 0 {
		util.ResponseError(c, "无效的岗位ID")
		return
	}

	if err := job.Delete(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}

	util.ResponseSuccess(c, gin.H{"message": "删除成功"})
}
