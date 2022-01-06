package main

import (
	"fmt"
	"whalego/database"
	"whalego/models/Channel"
)

func main() {
	db := database.Connect()

	models := map[string]interface{} {
		"Channel": Channel.New(),
	}

	for name, model := range models {
		db.AutoMigrate(model)
		fmt.Println(name + " migrated")
	}
}