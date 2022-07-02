package usersettingmodel

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserSettingModel struct {
	UserID       uint64
	BackGroupUrl string
}

func GetUserSetting(userID uint64) *UserSettingModel {
	result := &UserSettingModel{}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return &UserSettingModel{}
}
