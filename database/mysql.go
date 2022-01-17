package database

import (
	"whalego/Config"
	"whalego/errorHandler"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	host := Config.Database["HOST"]
	port := Config.Database["PORT"]
	dbName := Config.Database["DBNAME"]
	username := Config.Database["USERNAME"]
	pass := Config.Database["PASS"]

	dsn := username + ":" + pass + "@tcp(" + host + ":3306" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	errorHandler.LogFile(err)

	return db
}

func Close(connection *gorm.DB) {
	dbC, err := connection.DB()

	errorHandler.LogFile(err)

	dbC.Close()
}
