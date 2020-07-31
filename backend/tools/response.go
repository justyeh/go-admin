package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(c *gin.Context, data map[string]interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   data,
	})
}

func ResponseError(c *gin.Context, status int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": message,
	})
}
