package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var RootPath string

func InitConfig(path string) error {
	RootPath = path

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	return viper.ReadInConfig()
}

func GetMysqlDsn() string {
	ip := viper.GetString("mysql.ip")
	port := viper.GetInt("mysql.port")
	userName := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	database := viper.GetString("mysql.database")
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userName, password, ip, port, database)
}

func GetString(key string) string {
	return viper.GetString(key)
}
