package utils

import (
	"github.com/gin-gonic/gin"
)

func HandleSuccess(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"result": data,
	})
}

func HandleError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}