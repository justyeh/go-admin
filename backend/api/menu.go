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
		"list": sliceToTree(list),
	})
}

func AddMenu(c *gin.Context) {
	menu := models.Menu{ID: tools.UUID()}
	if err := c.ShouldBind(&menu); err != nil {
		tools.ResponseBindError(c, err)
		return
	}
	if err := menu.Create(); err != nil {
		tools.ResponseError(c, err.Error())
	}
	tools.ResponseSuccess(c, gin.H{"data": menu})
}

func EditMenu(c *gin.Context) {

}

func DeleteMenu(c *gin.Context) {

}

func sliceToTree(source []models.Menu) []models.Menu {
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
		handleChildNode(&result, item)
	}

	return result
}

func handleChildNode(list *[]models.Menu, m models.Menu) {
	for index, item := range *list {
		if item.ID == m.Pid {
			(*list)[index].Children = append(item.Children, m)
			goto END
		}

		if len(item.Children) > 0 {
			handleChildNode(&(*list)[index].Children, m)
		}
	}
END:
	return
}
