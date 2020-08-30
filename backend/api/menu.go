package api

import (
	"backend/models"
	"backend/tools"

	"github.com/gin-gonic/gin"
)

func existInSlice(source []int, target int) bool {
	for _, item := range source {
		if item == target {
			return true
		}
	}
	return false
}

func sliceToTree(source []Menu) []Menu {
	result := []Menu{}

	// 获取id集合
	ids := []int{}
	for _, item := range source {
		ids = append(ids, item.ID)
	}

	// 遍历，找到所有根节点
	sourceCopy := make([]Menu, len(source))
	copy(sourceCopy, source)

	for index, item := range sourceCopy {
		if !existInSlice(ids, item.Pid) {
			resultLen := len(result)
			source = append(source[:index-resultLen], source[index-resultLen+1:]...)

			result = append(result, item)
		}
	}

	// 遍历，处理所有子节点
	for _, item := range source {
		HandleChild(&result, item)
	}

	return result
}

func HandleChild(list *[]Menu, m Menu) {
	for index, item := range *list {
		if item.ID == m.Pid {
			(*list)[index].Children = append(item.Children, m)
			goto END
		}

		if len(item.Children) > 0 {
			HandleChild(&(*list)[index].Children, m)
		}
	}
END:
	return
}

func MenuTree(c *gin.Context) {
	list, err := models.MenuTreeWithName(c.Query("keyword"))
	if err != nil {
		tools.ResponseError(c, err.Error())
		return
	}

	// list := tools.SliceToTree(list)

	tools.ResponseSuccess(c, gin.H{
		"list": list,
	})
}

func AddMenu(c *gin.Context) {

}

func EditMenu(c *gin.Context) {

}

func DeleteMenu(c *gin.Context) {

}
