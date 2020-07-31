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

}

func EditUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
