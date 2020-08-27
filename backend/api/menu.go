package api

import (
	"backend/models"
	"backend/tools"

	"github.com/gin-gonic/gin"
)

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
