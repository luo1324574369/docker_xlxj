package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"miniprogram/util/config"
)

func GetDB() (*gorm.DB, error) {
	dsn := config.GetMysqlDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
