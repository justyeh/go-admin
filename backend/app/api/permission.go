package app

import (
	"backend/models"
	"backend/util"

	"github.com/gin-gonic/gin"
)

func PermissionTree(c *gin.Context) {
	permission := models.Permission{Name: c.Query("keyword")}
	list, err := permission.PermissionTree()
	if err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{
		"list": permissionSliceToTree(list),
	})
}

func AddPermission(c *gin.Context) {
	now := util.GetUnixNow()
	permission := models.Permission{CreateAt: now, UpdateAt: now}
	if err := c.ShouldBind(&permission); err != nil {
		util.ResponseBindError(c, err)
		return
	}

	permission.ID = util.UUID()

	if err := permission.Create(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "添加成功", "data": permission})
}

func EditPermission(c *gin.Context) {
	permission := models.Permission{UpdateAt: util.GetUnixNow()}
	if err := c.ShouldBind(&permission); err != nil {
		util.ResponseBindError(c, err)
		return
	}
	childIds := util.GetChildIds("permission", permission.ID)
	if util.IsExistInStringSlice(childIds, permission.Pid) || permission.Pid == permission.ID {
		util.ResponseError(c, "参数不合法，pid不能为其本身或后代id")
		return
	}
	if err := permission.Update(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func DeletePermission(c *gin.Context) {
	permission := models.Permission{ID: c.Param("id")}

	childIds := util.GetChildIds("permission", permission.ID)
	ids := []string{permission.ID}
	ids = append(ids, childIds...)

	if err := permission.Delete(ids); err != nil {
		util.ResponseError(c, err.Error())
		return
	}

	util.ResponseSuccess(c, gin.H{"message": "删除成功"})
}

func permissionSliceToTree(permissionList []models.Permission) []models.Permission {
	// 获取id集合
	ids := []interface{}{}
	for _, item := range permissionList {
		ids = append(ids, item.ID)
	}

	// 遍历，找到所有根节点、后代节点
	rootNodes := []models.Permission{}
	childNodes := []models.Permission{}
	for _, item := range permissionList {
		if !util.IsExistInSlice(ids, item.Pid) {
			rootNodes = append(rootNodes, item)
		} else {
			childNodes = append(childNodes, item)
		}
	}

	handlePermissionNodeRelation(&childNodes, &rootNodes)
	return rootNodes
}

func handlePermissionNodeRelation(childNodes, parentNodes *[]models.Permission) {
	// 理论最多执行 n+n-1+n-2+...+1 次，即每次最后一个处理成功
	maxExectionTimes := (1 + len(*childNodes)) * len(*childNodes) / 2

	for len(*childNodes) > 0 && maxExectionTimes > 0 {
		for cIndex, child := range *childNodes {
			IS_PERMISSION_INSERT_SUCCESS = false
			permissionRecursive(parentNodes, child)
			if IS_PERMISSION_INSERT_SUCCESS {
				*childNodes = append((*childNodes)[:cIndex], (*childNodes)[cIndex+1:]...)
				break
			}
		}
		maxExectionTimes--
	}
}

var IS_PERMISSION_INSERT_SUCCESS = false

func permissionRecursive(list *[]models.Permission, target models.Permission) {
	for index, item := range *list {
		if item.ID == target.Pid {
			(*list)[index].Children = append(item.Children, target)
			IS_PERMISSION_INSERT_SUCCESS = true
			return
		} else if len(item.Children) > 0 {
			permissionRecursive(&(*list)[index].Children, target)
		}
	}
}
