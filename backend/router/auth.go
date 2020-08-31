package router

import (
	"backend/api"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRouter(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", api.Login)
		auth.POST("/logout", api.Logout)
		auth.POST("/updatePassword", api.UpdatePassword)
		auth.GET("/captcha", api.Captcha)
	}
}
