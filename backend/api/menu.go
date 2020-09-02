package api

import (
	"backend/models"
	"backend/tools"

	"github.com/gin-gonic/gin"
)

func MenuTree(c *gin.Context) {
	menu := models.Menu{Name: c.Query("keyword")}
	list, err := menu.MenuTreeWithName()
	if err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{
		"list": menuSliceToTree(list),
	})
}

func AddMenu(c *gin.Context) {
	menu := models.Menu{}
	if err := c.ShouldBind(&menu); err != nil {
		tools.ResponseBindError(c, err)
		return
	}

	now := tools.GetUnixNow()
	menu.ID = tools.UUID()
	menu.CreateAt = now
	menu.UpdateAt = now

	if err := menu.Create(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "添加成功", "data": menu})
}

func EditMenu(c *gin.Context) {
	menu := models.Menu{}
	if err := c.ShouldBind(&menu); err != nil {
		tools.ResponseBindError(c, err)
		return
	}

	menu.UpdateAt = tools.GetUnixNow()

	if err := menu.Update(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}
	tools.ResponseSuccess(c, gin.H{"message": "修改成功"})
}

func DeleteMenu(c *gin.Context) {
	menu := models.Menu{ID: c.Param("id")}

	if len(menu.ID) == 0 {
		tools.ResponseError(c, "无效的菜单ID")
		return
	}

	if err := menu.Delete(); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}

	tools.ResponseSuccess(c, gin.H{"message": "删除成功"})
}

func menuSliceToTree(source []models.Menu) []models.Menu {
	result := []models.Menu{}

	// 获取id集合
	ids := []interface{}{}
	for _, item := range source {
		ids = append(ids, item.ID)
	}

	// 遍历，找到所有根节点
	sourceCopy := make([]models.Menu, len(source))
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
		handleMenuChildNode(&result, item)
	}

	return result
}

func handleMenuChildNode(list *[]models.Menu, m models.Menu) {
	for index, item := range *list {
		if item.ID == m.Pid {
			(*list)[index].Children = append(item.Children, m)
			goto END
		}

		if len(item.Children) > 0 {
			handleMenuChildNode(&(*list)[index].Children, m)
		}
	}
END:
	return
}
