package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/whaleproxy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// TODO: use err handler here
	if err != nil {
		panic(err)
	}

	return db
}