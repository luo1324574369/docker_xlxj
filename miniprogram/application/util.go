package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func responseOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 2,
		"msg":  "",
		"data": nil,
	})
}

func responseErr(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  err.Error(),
		"data": nil,
	})
}

func responseNotLogin(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "not login",
	})
}

func response(c *gin.Context, code int, msg string, data map[string]interface{}) {
	c.JSON(code, gin.H{
		"code": 2,
		"msg":  msg,
		"data": data,
	})
}
