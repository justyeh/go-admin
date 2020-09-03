package api

import (
	"backend/models"
	"backend/tools"

	"github.com/gin-gonic/gin"
)

func JobList(c *gin.Context) {
	job := models.Job{Name: c.Query("keyword")}
	page := tools.NewPagination(c)

	list, total, err := job.JobList(page)
	if err != nil {
		tools.ResponseError(c, err.Error())
		return
	}

	tools.ResponseSuccess(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func AddJob(c *gin.Context) {
	job := models.Job{}
	if err := c.ShouldBind(&job); err != nil {
		tools.ResponseBindError(c, err)
		return
	}

	now := tools.GetUnixNow()
	job.ID = tools.UUID()
	job.CreateAt = now
	job.UpdateAt = now

	if err := job.Create(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "添加成功", "data": job})
}

func EditJob(c *gin.Context) {
	job := models.Job{}
	if err := c.ShouldBind(&job); err != nil {
		tools.ResponseBindError(c, err)
		return
	}

	job.UpdateAt = tools.GetUnixNow()

	if err := job.Update(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func DeleteJob(c *gin.Context) {
	job := models.Job{ID: c.Param("id")}

	if len(job.ID) == 0 {
		tools.ResponseError(c, "无效的岗位ID")
		return
	}

	if err := job.Delete(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}

	tools.ResponseSuccess(c, gin.H{"message": "删除成功"})
}
