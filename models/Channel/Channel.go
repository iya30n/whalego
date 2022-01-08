package Channel

import (
	"os"
	"whalego/database"
	"whalego/errorHandler"
	"whalego/services/telegram/ChatService"

	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model
	Username             string `gorm: "NOT NULL;size:256"`
	ChatId               int64  `gorm: "int, unique;"`
	Last_message_receive int64    `gorm: "int;"`
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

func (c *Channel) Update(data map[string]interface{}) {
	db := database.Connect()

	db.Model(&c).Updates(data)
}

func (c *Channel) Delete() {
	db := database.Connect()

	db.Unscoped().Delete(&c)
}

func (c *Channel) GetChatId() int64 {
	if c.ChatId != 0 {
		return c.ChatId
	}

	chat, err := ChatService.New().GetChatId(c.Username)

	 if err.Error() == "USERNAME_NOT_OCCUPIED" && chat == nil {
		c.Delete()
		os.Exit(1)
	}

	errorHandler.LogFile(err)

	c.Update(map[string]interface{} {
		"chat_id": chat.Id,
	})

	return chat.Id
}