package logic

import (
	"miniprogram/model/usersettingmodel"
)

type User struct {
	UserID uint64
}

func (u *User) SetBgUrl(bgUrl string) error {
	userSetting, err := usersettingmodel.GetUserSetting(u.UserID)
	if err != nil {
		return err
	}
	if userSetting.UserID == 0 {
		data := &usersettingmodel.UserSettingModel{UserID: u.UserID, BgUrl: bgUrl}
		return usersettingmodel.CreateUserSetting(data)
	}
	userSetting.BgUrl = bgUrl
	return usersettingmodel.UpdateUserSetting(userSetting)
}

func (u *User) GetBgUrl() string {
	userSetting, err := usersettingmodel.GetUserSetting(u.UserID)
	if err != nil {
		return ""
	}
	return userSetting.BgUrl
}
