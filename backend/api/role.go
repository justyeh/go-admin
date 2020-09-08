package api

import (
	"backend/models"
	"backend/tools"

	"github.com/gin-gonic/gin"
)

func RoleList(c *gin.Context) {
	role := models.Role{Name: c.Query("keyword")}
	page := tools.NewPagination(c)

	list, count, err := role.RoleList(page)
	if err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{
		"list":  list,
		"total": count,
	})
}

func AddRole(c *gin.Context) {
	now := tools.GetUnixNow()
	role := models.Role{CreateAt: now, UpdateAt: now}
	if err := c.ShouldBind(&role); err != nil {
		tools.ResponseBindError(c, err)
		return
	}
	role.ID = tools.UUID()
	if err := role.Create(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "添加成功", "data": role})
}

func EditRole(c *gin.Context) {
	role := models.Role{UpdateAt: tools.GetUnixNow()}
	if err := c.ShouldBind(&role); err != nil {
		tools.ResponseBindError(c, err)
		return
	}
	if err := role.Update(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func UpdateRoleStatus(c *gin.Context) {
	role := models.Role{Name: "temp", UpdateAt: tools.GetUnixNow()}

	if err := c.ShouldBind(&role); err != nil {
		tools.ResponseBindError(c, err)
		return
	}

	if len(role.ID) == 0 {
		tools.ResponseError(c, "无效的角色ID")
		return
	}

	if err := role.UpdateStatus(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func DeleteRole(c *gin.Context) {
	role := models.Role{ID: c.Param("id")}

	if len(role.ID) == 0 {
		tools.ResponseError(c, "无效的角色ID")
		return
	}

	if err := role.Delete(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}

	tools.ResponseSuccess(c, gin.H{"message": "删除成功"})
}

func RoleMenuIds(c *gin.Context) {
	role := models.Role{ID: c.Param("roleId")}
	if len(role.ID) == 0 {
		tools.ResponseError(c, "无效的角色Id")
		return
	}
	ids, err := role.RoleMenuIds()
	if err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{
		"ids": ids,
	})
}

func RolePermissionIds(c *gin.Context) {
	role := models.Role{ID: c.Param("roleId")}
	if len(role.ID) == 0 {
		tools.ResponseError(c, "无效的角色Id")
		return
	}
	ids, err := role.RolePermissionIds()
	if err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{
		"ids": ids,
	})
}

func UpdateRoleMenu(c *gin.Context) {
	roleMenu := models.RoleMenu{}
	if err := c.ShouldBind(&roleMenu); err != nil {
		tools.ResponseBindError(c, err)
		return
	}
	if err := roleMenu.UpdateRoleMenu(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "操作成功"})
}

func UpdateRolePermission(c *gin.Context) {
	rolePermission := models.RolePermission{}
	if err := c.ShouldBind(&rolePermission); err != nil {
		tools.ResponseBindError(c, err)
		return
	}
	if err := rolePermission.UpdateRolePermission(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "操作成功"})
}
