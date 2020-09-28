package app

import (
	"backend/models"
	"backend/util"

	"github.com/gin-gonic/gin"
)

func DeptTree(c *gin.Context) {
	dept := models.Dept{Name: c.Query("keyword")}
	list, err := dept.DeptTree()
	if err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{
		"list": deptSliceToTree(list),
	})
}

func AddDept(c *gin.Context) {
	now := util.GetUnixNow()
	dept := models.Dept{CreateAt: now, UpdateAt: now}
	if err := c.ShouldBind(&dept); err != nil {
		util.ResponseBindError(c, err)
		return
	}

	dept.ID = util.UUID()

	if err := dept.Create(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "添加成功", "data": dept})
}

func EditDept(c *gin.Context) {
	dept := models.Dept{UpdateAt: util.GetUnixNow()}
	if err := c.ShouldBind(&dept); err != nil {
		util.ResponseBindError(c, err)
		return
	}
	childIds := util.GetChildIds("dept", dept.ID)
	if util.IsExistInStringSlice(childIds, dept.Pid) || dept.Pid == dept.ID {
		util.ResponseError(c, "参数不合法，pid不能为其本身或后代id")
		return
	}
	if err := dept.Update(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func DeleteDept(c *gin.Context) {
	dept := models.Dept{ID: c.Param("id")}
	if len(dept.ID) == 0 {
		util.ResponseError(c, "无效的部门ID")
		return
	}

	childIds := util.GetChildIds("dept", dept.ID)
	ids := []string{dept.ID}
	ids = append(ids, childIds...)

	if err := dept.Delete(ids); err != nil {
		util.ResponseError(c, err.Error())
		return
	}

	util.ResponseSuccess(c, gin.H{"message": "删除成功"})
}

func deptSliceToTree(deptList []models.Dept) []models.Dept {
	// 获取id集合
	ids := []interface{}{}
	for _, item := range deptList {
		ids = append(ids, item.ID)
	}

	// 遍历，找到所有根节点、后代节点
	rootNodes := []models.Dept{}
	childNodes := []models.Dept{}
	for _, item := range deptList {
		if !util.IsExistInSlice(ids, item.Pid) {
			rootNodes = append(rootNodes, item)
		} else {
			childNodes = append(childNodes, item)
		}
	}

	handleDeptNodeRelation(&childNodes, &rootNodes)
	return rootNodes
}

func handleDeptNodeRelation(childNodes, parentNodes *[]models.Dept) {
	// 理论最多执行 n+n-1+n-2+...+1 次，即每次最后一个处理成功
	maxExectionTimes := (1 + len(*childNodes)) * len(*childNodes) / 2
	for len(*childNodes) > 0 && maxExectionTimes > 0 {
		for cIndex, child := range *childNodes {
			IS_DEPT_INSERT_SUCCESS = false
			deptRecursive(parentNodes, child)
			if IS_DEPT_INSERT_SUCCESS {
				*childNodes = append((*childNodes)[:cIndex], (*childNodes)[cIndex+1:]...)
				break
			}
		}
		maxExectionTimes--
	}
}

var IS_DEPT_INSERT_SUCCESS = false

func deptRecursive(list *[]models.Dept, target models.Dept) {
	for index, item := range *list {
		if item.ID == target.Pid {
			(*list)[index].Children = append(item.Children, target)
			IS_DEPT_INSERT_SUCCESS = true
			return
		} else if len(item.Children) > 0 {
			deptRecursive(&(*list)[index].Children, target)
		}
	}
}
