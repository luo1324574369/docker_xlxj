package logic

import "miniprogram/model/usersettingmodel"

type User struct {
	UserID uint64
}

func (u *User) GetBackGroupUrl() string {
	userSetting := usersettingmodel.GetUserSetting(u.UserID)
	return userSetting.BackGroupUrl
}
