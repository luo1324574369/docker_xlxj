package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello 9999",
		})
	})
	if r.Run() != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		return
	}
}
