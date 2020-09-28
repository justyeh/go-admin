package api

import (
	"backend/models"
	"backend/util"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	user := models.ViewUser{User: models.User{Account: c.Query("keyword"), Nickname: c.Query("keyword")}, DeptId: c.Query("deptId")}
	page := util.NewPagination(c)
	list, total, err := user.UserList(page)
	if err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func AddUser(c *gin.Context) {
	now := util.GetUnixNow()
	user := models.SystemUser{User: models.User{CreateAt: now, UpdateAt: now}}
	if err := c.ShouldBind(&user); err != nil {
		util.ResponseBindError(c, err)
		return
	}
	user.ID = util.UUID()
	user.Password = util.GetMD5("123456")
	if err := user.Create(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "添加成功", "data": user})
}

func EditUser(c *gin.Context) {
	user := models.SystemUser{User: models.User{UpdateAt: util.GetUnixNow()}}
	if err := c.ShouldBind(&user); err != nil {
		util.ResponseBindError(c, err)
		return
	}
	if err := user.Update(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func UpdateUserStatus(c *gin.Context) {
	userStatus := models.UserStatus{UpdateAt: util.GetUnixNow()}
	if err := c.ShouldBind(&userStatus); err != nil {
		util.ResponseBindError(c, err)
		return
	}
	if err := userStatus.UpdateStatus(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func DeleteUser(c *gin.Context) {
	user := models.User{ID: c.Param("id")}
	if user.ID == "0" {
		util.ResponseError(c, "改用户无法删除")
		return
	}
	if len(user.ID) == 0 {
		util.ResponseError(c, "无效的用户ID")
		return
	}
	if err := user.Delete(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "删除成功"})
}
