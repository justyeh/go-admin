package router

import (
	"backend/api"

	"github.com/gin-gonic/gin"
)

func RegisterSystemRouter(router *gin.RouterGroup) {
	user := router.Group("/system/user")
	{
		user.GET("/list", api.UserList)
		user.POST("/", api.AddUser)
		user.PUT("/", api.EditUser)
		user.PUT("/updateUserStatus", api.UpdateUserStatus)
		user.DELETE("/:id", api.DeleteUser)
	}

	menu := router.Group("/system/menu")
	{
		menu.GET("/tree", api.MenuTree)
		menu.POST("/", api.AddMenu)
		menu.PUT("/", api.EditMenu)
		menu.DELETE("/:id", api.DeleteMenu)
	}

	permission := router.Group("/system/permission")
	{
		permission.GET("/tree", api.PermissionTree)
		permission.POST("/", api.AddPermission)
		permission.PUT("/", api.EditPermission)
		permission.DELETE("/:id", api.DeletePermission)
	}

	role := router.Group("/system/role")
	{
		role.GET("/list", api.RoleList)
		role.POST("/", api.AddRole)
		role.PUT("/", api.EditRole)
		role.PUT("/updateRoleStatus", api.UpdateRoleStatus)
		role.DELETE("/:id", api.DeleteRole)
	}

	dept := router.Group("/system/dept")
	{
		dept.GET("/tree", api.DeptTree)
		dept.POST("/", api.AddDept)
		dept.PUT("/", api.EditDept)
		dept.DELETE("/:id", api.DeleteDept)
	}

	job := router.Group("/system/job")
	{
		job.GET("/list", api.JobList)
		job.POST("/", api.AddJob)
		job.PUT("/", api.EditJob)
		job.DELETE("/:id", api.DeleteJob)
	}

	dictionary := router.Group("/system/dictionary")
	{
		dictionary.GET("/main/list", api.DictionaryList)
		dictionary.POST("/main/", api.AddDictionary)
		dictionary.PUT("/main/", api.EditDictionary)
		dictionary.DELETE("/main/:id", api.DeleteDictionary)

		dictionary.GET("/detail/list", api.DictionaryDetailList)
		dictionary.POST("/detail/", api.AddDictionaryDetail)
		dictionary.PUT("/detail/", api.EditDictionaryDetail)
		dictionary.DELETE("/detail/:id", api.DeleteDictionaryDetail)
	}
}
