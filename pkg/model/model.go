package model

import (
	"goblog/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error
	config := mysql.New(mysql.Config{
		DSN: "root:a909958300a@tcp(106.75.169.123:4040)/goblog?charset=utf8&parseTime=True&loc=Local"})
	DB, err = gorm.Open(config, &gorm.Config{})
	logger.LogError(err)
	return DB
}
