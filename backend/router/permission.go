package router

import (
	"backend/api"

	"github.com/gin-gonic/gin"
)

func RegisterPermissionRouter(router *gin.RouterGroup) {
	permission := router.Group("/permission")
	{
		permission.GET("/", api.PermissionList)
		permission.POST("/", api.AddPermission)
		permission.PUT("/", api.EditPermission)
		permission.DELETE("/", api.DeletePermission)
	}
}
