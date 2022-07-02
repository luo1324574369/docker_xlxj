package main

import (
	"miniprogram/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	Server = gin.Default()
	a      = &application.App{}
	u      = &application.User{}
)

func init() {
	Server.POST("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 测试
	Server.GET("ping", a.Pong)
	// 获取用户背景图
	Server.GET("getBgUrl", u.GetBackGroupUrl)
}
