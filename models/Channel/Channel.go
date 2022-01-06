package Channel

import (
	"whalego/database"
	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model
	Username             string `gorm: "NOT NULL;size:256"`
	ChatId               int64  `gorm: "int, unique;"`
	Last_message_receive int    `gorm: "int;"`
	Handler              string `gorm: "NOT NULL;size:30"`
}

func New() *Channel {
	return &Channel{}
}

func (c *Channel) All() []Channel {
	db := database.Connect()

	var channels []Channel

	db.Find(&channels)
	
	return channels
}

func (c *Channel) FindByUsername(username string) *Channel {
	db := database.Connect()

	db.Find(&c, "username = ?", username)

	if c.Username == "" {
		panic("the channel " + username + " does not exists.")
	}

	return c
}
