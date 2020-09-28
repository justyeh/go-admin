package api

import (
	"backend/models"
	"backend/util"

	"github.com/gin-gonic/gin"
)

func RoleList(c *gin.Context) {
	role := models.Role{Name: c.Query("keyword")}
	page := util.NewPagination(c)

	list, count, err := role.RoleList(page)
	if err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{
		"list":  list,
		"total": count,
	})
}

func AddRole(c *gin.Context) {
	now := util.GetUnixNow()
	role := models.Role{CreateAt: now, UpdateAt: now}
	if err := c.ShouldBind(&role); err != nil {
		util.ResponseBindError(c, err)
		return
	}
	role.ID = util.UUID()
	if err := role.Create(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "添加成功", "data": role})
}

func EditRole(c *gin.Context) {
	role := models.Role{UpdateAt: util.GetUnixNow()}
	if err := c.ShouldBind(&role); err != nil {
		util.ResponseBindError(c, err)
		return
	}
	if err := role.Update(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func UpdateRoleStatus(c *gin.Context) {
	roleStatus := models.RoleStatus{UpdateAt: util.GetUnixNow()}
	if err := c.ShouldBind(&roleStatus); err != nil {
		util.ResponseBindError(c, err)
		return
	}
	if err := roleStatus.UpdateStatus(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func DeleteRole(c *gin.Context) {
	role := models.Role{ID: c.Param("id")}

	if len(role.ID) == 0 {
		util.ResponseError(c, "无效的角色ID")
		return
	}

	if err := role.Delete(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}

	util.ResponseSuccess(c, gin.H{"message": "删除成功"})
}

func RoleMenuIds(c *gin.Context) {
	role := models.Role{ID: c.Param("roleId")}
	if len(role.ID) == 0 {
		util.ResponseError(c, "无效的角色Id")
		return
	}
	ids, err := role.RoleMenuIds()
	if err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{
		"ids": ids,
	})
}

func RolePermissionIds(c *gin.Context) {
	role := models.Role{ID: c.Param("roleId")}
	if len(role.ID) == 0 {
		util.ResponseError(c, "无效的角色Id")
		return
	}
	ids, err := role.RolePermissionIds()
	if err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{
		"ids": ids,
	})
}

func UpdateRoleMenu(c *gin.Context) {
	roleMenu := models.RoleMenu{}
	if err := c.ShouldBind(&roleMenu); err != nil {
		util.ResponseBindError(c, err)
		return
	}
	if err := roleMenu.UpdateRoleMenu(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "操作成功"})
}

func UpdateRolePermission(c *gin.Context) {
	rolePermission := models.RolePermission{}
	if err := c.ShouldBind(&rolePermission); err != nil {
		util.ResponseBindError(c, err)
		return
	}
	if err := rolePermission.UpdateRolePermission(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "操作成功"})
}
