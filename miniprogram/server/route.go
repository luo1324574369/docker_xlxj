package server

import (
	"miniprogram/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	a = &application.App{}
	u = &application.User{}
)

func setRoute(engine *gin.Engine) {
	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 测试
	engine.GET("ping", a.Pong)
	// 使用code获取openid和sessionKey
	engine.GET("wxLogin", a.WxLogin)

	authorized := engine.Group("/wxUser", WXUserInfoMiddleware())
	// 获取用户背景图
	authorized.GET("getBgUrl", u.GetBgUrl)
	// 设置用户背景图
	authorized.POST("setBgUrl", u.SetBgUrl)
	// 上传图片
	authorized.POST("uploadPicture", a.UploadPicture)
}
