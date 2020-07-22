package router

import (
	"backend/middleware"
	"backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//InitRouter  初始化路由
func InitRouter() *gin.Engine {
	router := gin.New()
	router.GET("/", home)

	wiki := router.Group("/wiki")
	wiki.GET("/", service.GetWikiList)

	router.NoRoute(notFound)
	middleware.InitMiddleware(router)

	router.Run(":1234")
	return router
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": "欢迎访问中锐滨湖尚城wiki",
	})
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": "404 not found",
	})
}
