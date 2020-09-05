package api

import (
	"backend/models"
	"backend/tools"

	"github.com/gin-gonic/gin"
)

func PermissionTree(c *gin.Context) {
	permission := models.Permission{Name: c.Query("keyword")}
	list, err := permission.PermissionTree()
	if err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{
		"list": permissionSliceToTree(list),
	})
}

func AddPermission(c *gin.Context) {
	now := tools.GetUnixNow()
	permission := models.Permission{CreateAt: now, UpdateAt: now}
	if err := c.ShouldBind(&permission); err != nil {
		tools.ResponseBindError(c, err)
		return
	}

	permission.ID = tools.UUID()

	if err := permission.Create(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "添加成功", "data": permission})
}

func EditPermission(c *gin.Context) {
	permission := models.Permission{UpdateAt: tools.GetUnixNow()}
	if err := c.ShouldBind(&permission); err != nil {
		tools.ResponseBindError(c, err)
		return
	}

	if permission.ID == permission.Pid {
		tools.ResponseError(c, "参数不合法，pid不能等于id")
		return
	}

	if err := permission.Update(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func DeletePermission(c *gin.Context) {
	permission := models.Permission{ID: c.Param("id")}

	if len(permission.ID) == 0 {
		tools.ResponseError(c, "无效的权限ID")
		return
	}

	if err := permission.Delete(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}

	tools.ResponseSuccess(c, gin.H{"message": "删除成功"})
}

func permissionSliceToTree(source []models.Permission) []models.Permission {
	result := []models.Permission{}

	// 获取id集合
	ids := []interface{}{}
	for _, item := range source {
		ids = append(ids, item.ID)
	}

	// 遍历，找到所有根节点
	sourceCopy := make([]models.Permission, len(source))
	copy(sourceCopy, source)

	for index, item := range sourceCopy {
		if !tools.IsExistInSlice(ids, item.Pid) {
			resultLen := len(result)
			source = append(source[:index-resultLen], source[index-resultLen+1:]...)

			result = append(result, item)
		}
	}

	// 遍历，处理所有子节点
	for _, item := range source {
		handlePermissionChildNode(&result, item)
	}

	return result
}

func handlePermissionChildNode(list *[]models.Permission, m models.Permission) {
	for index, item := range *list {
		if item.ID == m.Pid {
			(*list)[index].Children = append(item.Children, m)
			goto END
		}

		if len(item.Children) > 0 {
			handlePermissionChildNode(&(*list)[index].Children, m)
		}
	}
END:
	return
}
