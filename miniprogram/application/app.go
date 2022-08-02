package application

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"miniprogram/model/usermodel"
	"miniprogram/util"
	"miniprogram/util/config"
	"miniprogram/util/logger"
	"net/http"
	"net/url"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type App struct{}

func (a *App) Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "app pong",
	})
}

type wxLoginRet struct {
	ErrCode    json.Number `json:"errcode"`
	SessionKey string      `json:"session_key"`
	OpenID     string      `json:"openid"`
	ErrMsg     string      `json:"errmsg"`
}

func (a *App) WxLogin(c *gin.Context) {
	code := c.Query("code")
	appId := config.GetString("wx.app_id")
	secret := config.GetString("wx.secret")

	loginUrl := "https://api.weixin.qq.com/sns/jscode2session?" +
		"appid=" + appId + "&secret=" + secret + "&js_code=" + code + "&grant_type=authorization_code"
	client := &http.Client{}
	req, buildErr := http.NewRequest("GET", loginUrl, nil)
	if buildErr != nil {
		responseErr(c, buildErr)
		return
	}
	resp, doErr := client.Do(req)
	if doErr != nil {
		responseErr(c, doErr)
		return
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logger.NewLogger().Error(fmt.Sprintf("WxLogin err : %v", err))
		}
	}()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		responseErr(c, readErr)
		return
	}
	wxRet := wxLoginRet{}
	if err := json.Unmarshal(body, &wxRet); err != nil {
		responseErr(c, err)
		return
	}
	if wxRet.ErrCode != "" {
		responseErr(c, errors.New("err code"))
		return
	}

	// 存库
	sessionKey := fmt.Sprintf("%x", md5.Sum([]byte(wxRet.SessionKey)))
	data := &usermodel.UserModel{OpenID: wxRet.OpenID, SessionKey: sessionKey}
	if err := usermodel.CreateOrUpdateUser(data); err != nil {
		logger.NewLogger().Error(fmt.Sprintf("createUser get err %v", err))
		responseErr(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 2,
		"msg":  "",
		"data": map[string]string{
			"token": url.QueryEscape(wxRet.SessionKey),
		},
	})
}

func (a *App) UploadPicture(c *gin.Context) {
	file, err := c.FormFile("pic")
	if err != nil {
		responseErr(c, err)
		return
	}

	baseName := util.RandStringRunes(10) + path.Ext(file.Filename)
	basePath := config.GetString("upload.path")
	filename := basePath + filepath.Base(baseName)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		responseErr(c, err)
		return
	}
	response(c, http.StatusOK, "", map[string]interface{}{
		"file_name": filename,
	})
}
