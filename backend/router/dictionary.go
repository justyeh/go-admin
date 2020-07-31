package router

import (
	"backend/api"

	"github.com/gin-gonic/gin"
)

func RegisterDictionaryRouter(router *gin.RouterGroup) {
	dictionary := router.Group("/dictionary")
	{
		dictionary.GET("/", api.DictionaryList)
		dictionary.POST("/", api.AddDictionary)
		dictionary.PUT("/", api.EditDictionary)
		dictionary.DELETE("/", api.DeleteDictionary)
	}
}
