package initialize

import (
	"backend/middleware"
	"backend/router"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {

	gin.SetMode(CONFIG.Gin.Mode)
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

	fmt.Println("程序启动成功，服务运行与http://127.0.0.1:" + CONFIG.Gin.Port)
	r.Run(":" + CONFIG.Gin.Port)
}
