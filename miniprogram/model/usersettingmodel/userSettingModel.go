package usersettingmodel

import (
	"errors"
	"fmt"
	"miniprogram/model"
	"miniprogram/util/logger"
	"time"
)

type UserSettingModel struct {
	UserID     uint64
	BgUrl      string    `gorm:"column:bg_url"`
	CreateTime time.Time `gorm:"autoCreateTime:true"`
	ModifyTime time.Time `gorm:"autoUpdateTime:true"`
}

func (u UserSettingModel) TableName() string {
	return "t_user_setting"
}

func CreateUserSetting(userSetting *UserSettingModel) error {
	db, err := model.GetDB()
	if err != nil {
		return errors.New("db error")
	}
	if result := db.Create(userSetting); result.Error != nil {
		logger.NewLogger().Error(fmt.Sprintf("CreateUserSetting err %v", result.Error))
		return errors.New("insert err")
	}
	return nil
}

func UpdateUserSetting(userSetting *UserSettingModel) error {
	db, err := model.GetDB()
	if err != nil {
		return errors.New("db error")
	}

	if result := db.Where("user_id", userSetting.UserID).Save(userSetting); result.Error != nil {
		logger.NewLogger().Error(fmt.Sprintf("UpdateUserSetting err %v", result.Error))
		return errors.New("update err")
	}
	return nil
}

func GetUserSetting(userID uint64) (*UserSettingModel, error) {
	result := &UserSettingModel{}
	db, err := model.GetDB()
	if err != nil {
		return nil, errors.New("db error")
	}

	if db.Where("user_id = ?", userID).First(result); db.Error != nil {
		logger.NewLogger().Error(fmt.Sprintf("GetUserSetting err %v", db.Error))
		return nil, db.Error
	}
	return result, nil
}
