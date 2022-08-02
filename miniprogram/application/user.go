package application

import (
	"miniprogram/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct{}

// GetBgUrl 获取用户背景图片
func (u *User) GetBgUrl(c *gin.Context) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		responseNotLogin(c)
		return
	}

	user := &logic.User{UserID: uint64(userID)}
	response(c, http.StatusOK, "", map[string]interface{}{
		"bg_url": user.GetBgUrl(),
	})
}

// SetBgUrl 设置用户背景图片
func (u *User) SetBgUrl(c *gin.Context) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		responseNotLogin(c)
		return
	}

	BgUrl := c.PostForm("bg_url")
	user := &logic.User{UserID: userID}
	if err := user.SetBgUrl(BgUrl); err != nil {
		responseErr(c, err)
		return
	}
	responseOK(c)
}
