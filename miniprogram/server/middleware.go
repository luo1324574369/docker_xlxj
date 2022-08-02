package server

import (
	"crypto/md5"
	"fmt"
	"miniprogram/model/usermodel"
	"miniprogram/util/logger"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func WXUserInfoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, queryErr := url.QueryUnescape(c.GetHeader("token"))
		if token == "" || queryErr != nil {
			c.Redirect(http.StatusFound, "/")
		}

		sessionKey := fmt.Sprintf("%x", md5.Sum([]byte(token)))
		user, getErr := usermodel.GetUserBySessionKey(sessionKey)
		if getErr != nil {
			logger.NewLogger().Error(fmt.Sprintf("wxUserMiddleware err token %s", token))
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "not login",
			})
			return
		}

		if user.UserID == 0 {
			logger.NewLogger().Error(fmt.Sprintf("wxUserMiddleware err token %s", token))
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "not login",
			})
			return
		}

		c.Set("user_id", user.UserID)
		c.Next()
	}
}
