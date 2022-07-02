package application

import (
	"github.com/gin-gonic/gin"
)

func response(c *gin.Context, code int, msg string, data map[string]interface{}) {
	c.JSON(code, gin.H{
		"msg":  msg,
		"data": data,
	})
}
