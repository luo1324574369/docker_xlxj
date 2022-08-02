package usermodel

import (
	"errors"
	"fmt"
	"miniprogram/model"
	"miniprogram/util/logger"
	"time"
)

type UserModel struct {
	UserID     uint64 `gorm:"column:id"`
	OpenID     string
	SessionKey string
	CreateTime time.Time `gorm:"autoCreateTime:true"`
	ModifyTime time.Time `gorm:"autoUpdateTime:true"`
}

func (u UserModel) TableName() string {
	return "t_user"
}

func CreateOrUpdateUser(user *UserModel) error {
	db, err := model.GetDB()
	if err != nil {
		return errors.New("db error")
	}

	sql := fmt.Sprintf("INSERT INTO t_user (`open_id`,`session_key`) values ('%s','%s') ON DUPLICATE KEY UPDATE `session_key` = '%s'", user.OpenID, user.SessionKey, user.SessionKey)
	if result := db.Exec(sql); result.Error != nil {
		logger.NewLogger().Error(fmt.Sprintf("CreateOrUpdateUser err %v", result.Error))
		return errors.New("insert err")
	}
	return nil
}

func GetUserBySessionKey(sessionKey string) (*UserModel, error) {
	result := &UserModel{}
	db, err := model.GetDB()
	if err != nil {
		return nil, errors.New("db error")
	}

	if db.Where("session_key = ?", sessionKey).First(result); db.Error != nil {
		logger.NewLogger().Error(fmt.Sprintf("getUser err %v", db.Error))
		return nil, db.Error
	}
	return result, nil
}

func GetUserByOpenID(openID string) (*UserModel, error) {
	result := &UserModel{}
	db, err := model.GetDB()
	if err != nil {
		return nil, errors.New("db error")
	}

	if db.Where("open_id = ?", openID).First(result); db.Error != nil {
		logger.NewLogger().Error(fmt.Sprintf("getUser err %v", db.Error))
		return nil, db.Error
	}
	return result, nil
}
