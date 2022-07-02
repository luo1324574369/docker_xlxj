package application

import (
	"net/http"
	"strconv"

	"miniprogram/logic"

	"github.com/gin-gonic/gin"
)

type User struct{}

// GetBackGroupUrl 获取用户背景图片
func (u *User) GetBackGroupUrl(c *gin.Context) {
	paramUserID := c.Query("user_id")
	userID, err := strconv.Atoi(paramUserID)
	if err != nil {
		response(c, http.StatusOK, "参数错误", nil)
		return
	}

	user := &logic.User{UserID: uint64(userID)}
	response(c, http.StatusOK, "", map[string]interface{}{
		"bg_url": user.GetBackGroupUrl(),
	})
}

func (u User) GetSour() {

}
