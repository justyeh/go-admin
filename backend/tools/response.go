package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(c *gin.Context, data gin.H) {
	data["status"] = http.StatusOK
	c.JSON(http.StatusOK, data)
}

func ResponseError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusBadRequest,
		"message": message,
	})
}
