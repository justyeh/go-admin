package router

import (
	"backend/api"

	"github.com/gin-gonic/gin"
)

func RegisterMenuRouter(router *gin.RouterGroup) {
	menu := router.Group("/menu")
	{
		menu.GET("/", api.MenuList)
		menu.POST("/", api.AddMenu)
		menu.PUT("/", api.EditMenu)
		menu.DELETE("/", api.DeleteMenu)
	}
}
