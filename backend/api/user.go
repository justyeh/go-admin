package api

import (
	"backend/model"
	"backend/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	list, err := model.User{}.List()
	if err != nil {
		tools.ResponseError(c, http.StatusInternalServerError, "获取用户列表失败："+err.Error())
		return
	}

	p := tools.NewPagination(c)
	tools.ResponseSuccess(c, gin.H{
		"list": list,
		"page": p,
	})
}

func AddUser(c *gin.Context) {
	user := model.User{}
	user.Name = c.Param("name")

	insertId, err := user.Create()

	if err != nil {
		tools.ResponseError(c, http.StatusInternalServerError, "添加用户失败："+err.Error())
		return
	}
	user.ID = insertId

	return tools.ResponseSuccess(c, gin.H{"user": user})

}

func EditUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
