package database

import (
	"whalego/errorHandler"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/whaleproxy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	errorHandler.LogFile(err)

	return db
}

func Close(connection *gorm.DB) {
	dbC, err := connection.DB()

	errorHandler.LogFile(err)

	dbC.Close()
}