package service

import (
	"backend/model"
	"backend/tools"

	"github.com/gin-gonic/gin"
)

//GetWikiList 获取wiki列表
func GetWikiList(c *gin.Context) {
	var wiki model.Wiki
	wiki.List()
	// result, err := wiki.List()
}

//AddWiki 添加wiki
func AddWiki(c *gin.Context) {
	var wiki model.Wiki
	wiki.Title = c.PostForm("title")
	wiki.Title = c.PostForm("content")
	wiki.Add()
}

//EditWiki 编辑wiki
func EditWiki(c *gin.Context) {
	var wiki model.Wiki
	wiki.Title = c.PostForm("title")
	wiki.Title = c.PostForm("content")
	wiki.Edit()
}

//DeleteWiki 删除wiki
func DeleteWiki(c *gin.Context) {
	var wiki model.Wiki
	wiki.ID = tools.StringToInt(c.Param("id"))
	wiki.Delete()
}
