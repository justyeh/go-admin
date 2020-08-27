package router

import (
	"backend/api"

	"github.com/gin-gonic/gin"
)

func RegisterSystemRouter(router *gin.RouterGroup) {
	menu := router.Group("/system/menu")
	{
		menu.GET("/tree", api.MenuTree)
		menu.POST("/", api.AddMenu)
		menu.PUT("/", api.EditMenu)
		menu.DELETE("/", api.DeleteMenu)
	}

	/* user := router.Group("/user")
	{
		user.GET("/", api.UserList)
		user.POST("/", api.AddUser)
		user.PUT("/", api.EditUser)
		user.DELETE("/", api.DeleteUser)
	}

	role := router.Group("/role")
	{
		role.GET("/", api.RoleList)
		role.POST("/", api.AddRole)
		role.PUT("/", api.EditRole)
		role.DELETE("/", api.DeleteRole)
	}

	permission := router.Group("/permission")
	{
		permission.GET("/", api.PermissionList)
		permission.POST("/", api.AddPermission)
		permission.PUT("/", api.EditPermission)
		permission.DELETE("/", api.DeletePermission)
	}

	dictionary := router.Group("/dictionary")
	{
		dictionary.GET("/", api.DictionaryList)
		dictionary.POST("/", api.AddDictionary)
		dictionary.PUT("/", api.EditDictionary)
		dictionary.DELETE("/", api.DeleteDictionary)
	} */
}
