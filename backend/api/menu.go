package api

import (
	"backend/models"
	"backend/util"

	"github.com/gin-gonic/gin"
)

func MenuTree(c *gin.Context) {
	menu := models.Menu{Name: c.Query("keyword")}
	list, err := menu.MenuTree()
	if err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{
		"list": menuSliceToTree(list),
	})
}

func AddMenu(c *gin.Context) {
	now := util.GetUnixNow()
	menu := models.Menu{CreateAt: now, UpdateAt: now}
	if err := c.ShouldBind(&menu); err != nil {
		util.ResponseBindError(c, err)
		return
	}

	menu.ID = util.UUID()

	if err := menu.Create(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "添加成功", "data": menu})
}

func EditMenu(c *gin.Context) {
	menu := models.Menu{UpdateAt: util.GetUnixNow()}
	if err := c.ShouldBind(&menu); err != nil {
		util.ResponseBindError(c, err)
		return
	}
	childIds := util.GetChildIds("menu", menu.ID)
	if util.IsExistInStringSlice(childIds, menu.Pid) || menu.Pid == menu.ID {
		util.ResponseError(c, "参数不合法，pid不能为其本身或后代id")
		return
	}
	if err := menu.Update(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func DeleteMenu(c *gin.Context) {
	menu := models.Menu{ID: c.Param("id")}
	if len(menu.ID) == 0 {
		util.ResponseError(c, "无效的菜单ID")
		return
	}

	childIds := util.GetChildIds("menu", menu.ID)
	ids := []string{menu.ID}
	ids = append(ids, childIds...)

	if err := menu.Delete(ids); err != nil {
		util.ResponseError(c, err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{"message": "删除成功"})
}

func menuSliceToTree(menuList []models.Menu) []models.Menu {
	// 获取id集合
	ids := []interface{}{}
	for _, item := range menuList {
		ids = append(ids, item.ID)
	}

	// 遍历，找到所有根节点、后代节点
	rootNodes := []models.Menu{}
	childNodes := []models.Menu{}
	for _, item := range menuList {
		if !util.IsExistInSlice(ids, item.Pid) {
			rootNodes = append(rootNodes, item)
		} else {
			childNodes = append(childNodes, item)
		}
	}

	handleMenuNodeRelation(&childNodes, &rootNodes)
	return rootNodes
}

func handleMenuNodeRelation(childNodes, parentNodes *[]models.Menu) {
	// 理论最多执行 n+n-1+n-2+...+1 次，即每次最后一个处理成功
	maxExectionTimes := (1 + len(*childNodes)) * len(*childNodes) / 2
	for len(*childNodes) > 0 && maxExectionTimes > 0 {
		for cIndex, child := range *childNodes {
			IS_MENU_INSERT_SUCCESS = false
			menuRecursive(parentNodes, child)
			if IS_MENU_INSERT_SUCCESS {
				*childNodes = append((*childNodes)[:cIndex], (*childNodes)[cIndex+1:]...)
				break
			}
		}
		maxExectionTimes--
	}
}

var IS_MENU_INSERT_SUCCESS = false

func menuRecursive(list *[]models.Menu, target models.Menu) {
	for index, item := range *list {
		if item.ID == target.Pid {
			(*list)[index].Children = append(item.Children, target)
			IS_MENU_INSERT_SUCCESS = true
			return
		} else if len(item.Children) > 0 {
			menuRecursive(&(*list)[index].Children, target)
		}
	}
}
