package api

import (
	"backend/models"
	"backend/tools"
	"fmt"

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
	childIds := tools.GetChildIds("dept", dept.ID)
	if tools.IsExistInStringSlice(childIds, dept.Pid) || dept.Pid == dept.ID {
		tools.ResponseError(c, "参数不合法，pid不能为其本身或后代id")
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

	childIds := tools.GetChildIds("dept", dept.ID)
	ids := []string{dept.ID}
	ids = append(ids, childIds...)

	if err := dept.Delete(ids); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}

	tools.ResponseSuccess(c, gin.H{"message": "删除成功"})
}

func deptSliceToTree(deptList []models.Dept) []models.Dept {
	result := []models.Dept{}

	// 获取id集合
	ids := []interface{}{}
	for _, item := range source {
		ids = append(ids, item.ID)
	}

	// 遍历，找到所有根节点
	sourceCopy := make([]models.Dept, len(source))
	for index, item := range sourceCopy {
		if !tools.IsExistInSlice(ids, item.Pid) {
			resultLen := len(result)
			source = append(source[:index-resultLen], source[index-resultLen+1:]...)
			result = append(result, item)
		}
	}

	fmt.Println(result)

	// 遍历，处理所有子节点
	handleDeptChildNode(&result, source)

	return result
}

func handleDeptChildNode(resultTree *[]models.Dept, waitHandleNodes []) {
	/* for len(*list) > 0 {
		for index, item := range *list {
			if item.ID == m.Pid {
				(*list)[index].Children = append(item.Children, m)
				*list = append((*list)[:index], (*list)[index+1:]...)
				goto END_THIS
			}

			if len(item.Children) > 0 {
				handleDeptChildNode(&(*list)[index].Children, m)
			}
		}
	END_THIS:
	}
	return */
}
