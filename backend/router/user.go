package router

import (
	"backend/api"

	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	{
		user.GET("/", api.UserList)
		user.POST("/", api.AddUser)
		user.PUT("/", api.EditUser)
		user.DELETE("/", api.DeleteUser)
	}
}
