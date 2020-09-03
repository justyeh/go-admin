package api

import (
	"backend/models"
	"backend/tools"

	"github.com/gin-gonic/gin"
)

func RoleList(c *gin.Context) {
	role := models.Role{Name: c.Query("keyword")}
	list, err := role.RoleList()
	if err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{
		"list": roleSliceToTree(list),
	})
}

func AddRole(c *gin.Context) {
	role := models.Role{}
	if err := c.ShouldBind(&role); err != nil {
		tools.ResponseBindError(c, err)
		return
	}

	now := tools.GetUnixNow()
	role.ID = tools.UUID()
	role.CreateAt = now
	role.UpdateAt = now

	if err := role.Create(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "添加成功", "data": role})
}

func EditRole(c *gin.Context) {
	role := models.Role{}
	if err := c.ShouldBind(&role); err != nil {
		tools.ResponseBindError(c, err)
		return
	}

	role.UpdateAt = tools.GetUnixNow()

	if err := role.Update(); err != nil {
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

func roleSliceToTree(source []models.Role) []models.Role {
	result := []models.Role{}

	// 获取id集合
	ids := []interface{}{}
	for _, item := range source {
		ids = append(ids, item.ID)
	}

	// 遍历，找到所有根节点
	sourceCopy := make([]models.Role, len(source))
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
		handleRoleChildNode(&result, item)
	}

	return result
}

func handleRoleChildNode(list *[]models.Role, m models.Role) {
	for index, item := range *list {
		if item.ID == m.Pid {
			(*list)[index].Children = append(item.Children, m)
			goto END
		}

		if len(item.Children) > 0 {
			handleRoleChildNode(&(*list)[index].Children, m)
		}
	}
END:
	return
}
