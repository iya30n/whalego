package database

import (
	"sync"
	"whalego/Config"
	"whalego/errorHandler"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var doneOnce sync.Once
var singletonConnection *gorm.DB

func Connect() *gorm.DB {
	doneOnce.Do(func() {
		config := Config.Get()

		host := config.Database.Host
		port := config.Database.Port
		dbName := config.Database.DBName
		username := config.Database.Username
		pass := config.Database.Password

		dsn := username + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		singletonConnection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		errorHandler.LogFile(err)
	})

	return singletonConnection
}

func Close(connection *gorm.DB) {
	/* dbC, err := connection.DB()

	errorHandler.LogFile(err)

	dbC.Close() */
}