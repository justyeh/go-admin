package router

import (
	"backend/api"

	"github.com/gin-gonic/gin"
)

func RegisterRoleRouter(router *gin.RouterGroup) {
	role := router.Group("/role")
	{
		role.GET("/", api.RoleList)
		role.POST("/", api.AddRole)
		role.PUT("/", api.EditRole)
		role.DELETE("/", api.DeleteRole)
	}
}
