package api

import (
	"backend/models"
	"backend/tools"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	user := models.ViewUser{User: models.User{Account: c.Query("keyword"), Nickname: c.Query("keyword")}, DeptId: c.Query("deptId")}
	page := tools.NewPagination(c)
	list, total, err := user.UserList(page)
	if err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func AddUser(c *gin.Context) {
	now := tools.GetUnixNow()
	user := models.SystemUser{User: models.User{CreateAt: now, UpdateAt: now}}
	if err := c.ShouldBind(&user); err != nil {
		tools.ResponseBindError(c, err)
		return
	}
	user.ID = tools.UUID()
	user.Password = tools.GetMD5("123456")
	if err := user.Create(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "添加成功", "data": user})
}

func EditUser(c *gin.Context) {
	user := models.SystemUser{User: models.User{UpdateAt: tools.GetUnixNow()}}
	if err := c.ShouldBind(&user); err != nil {
		tools.ResponseBindError(c, err)
		return
	}
	if err := user.Update(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func UpdateUserStatus(c *gin.Context) {
	userStatus := models.UserStatus{UpdateAt: tools.GetUnixNow()}
	if err := c.ShouldBind(&userStatus); err != nil {
		tools.ResponseBindError(c, err)
		return
	}
	if err := userStatus.UpdateStatus(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func DeleteUser(c *gin.Context) {
	user := models.User{ID: c.Param("id")}
	if user.ID == "0" {
		tools.ResponseError(c, "改用户无法删除")
		return
	}
	if len(user.ID) == 0 {
		tools.ResponseError(c, "无效的用户ID")
		return
	}
	if err := user.Delete(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "删除成功"})
}
