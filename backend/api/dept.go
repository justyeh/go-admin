package api

import (
	"backend/models"
	"backend/tools"

	"github.com/gin-gonic/gin"
)

func DeptTree(c *gin.Context) {
	dept := models.Dept{Name: c.Query("keyword")}
	list, err := dept.DeptTree()
	if err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{
		"list": deptSliceToTree(list),
	})
}

func AddDept(c *gin.Context) {
	now := tools.GetUnixNow()
	dept := models.Dept{CreateAt: now, UpdateAt: now}
	if err := c.ShouldBind(&dept); err != nil {
		tools.ResponseBindError(c, err)
		return
	}

	dept.ID = tools.UUID()

	if err := dept.Create(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "添加成功", "data": dept})
}

func EditDept(c *gin.Context) {
	dept := models.Dept{UpdateAt: tools.GetUnixNow()}
	if err := c.ShouldBind(&dept); err != nil {
		tools.ResponseBindError(c, err)
		return
	}

	if err := dept.Update(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func DeleteDept(c *gin.Context) {
	dept := models.Dept{ID: c.Param("id")}

	if len(dept.ID) == 0 {
		tools.ResponseError(c, "无效的部门ID")
		return
	}

	if err := dept.Delete(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}

	tools.ResponseSuccess(c, gin.H{"message": "删除成功"})
}

func deptSliceToTree(source []models.Dept) []models.Dept {
	result := []models.Dept{}

	// 获取id集合
	ids := []interface{}{}
	for _, item := range source {
		ids = append(ids, item.ID)
	}

	// 遍历，找到所有根节点
	sourceCopy := make([]models.Dept, len(source))
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
		handleDeptChildNode(&result, item)
	}

	return result
}

func handleDeptChildNode(list *[]models.Dept, m models.Dept) {
	for index, item := range *list {
		if item.ID == m.Pid {
			(*list)[index].Children = append(item.Children, m)
			goto END
		}

		if len(item.Children) > 0 {
			handleDeptChildNode(&(*list)[index].Children, m)
		}
	}
END:
	return
}
