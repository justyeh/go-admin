package initialize

import (
	"backend/middleware"
	"backend/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// 全局中间件
	middleware.RegisterGlobalMiddleware(r)

	// 根路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "欢迎访问中锐滨湖尚城wiki",
		})
	})

	// 404处理
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "无效的url链接",
		})
	})

	// 业务路由
	ApiGroup := r.Group("/api")
	router.RegisterAuthRouter(ApiGroup)       // 登录、注册、验证码等认证相关
	router.RegisterUserRouter(ApiGroup)       // 用户
	router.RegisterRoleRouter(ApiGroup)       // 角色
	router.RegisterPermissionRouter(ApiGroup) // 权限
	router.RegisterMenuRouter(ApiGroup)       // 菜单
	router.RegisterDictionaryRouter(ApiGroup) // 字典

	r.Run(":1234")
	return r
}
