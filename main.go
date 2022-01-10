package main

import (
	"fmt"

	// "whalego/errorHandler"
	"whalego/models/Channel"
	"whalego/services/telegram/MessageService"

	"github.com/zelenin/go-tdlib/client"
	// "whalego/database/migration"
)

// type DeepMap map[string][]*DeepMap
type DeepMap struct {
	data map[string]DeepMap
}

func main() {
	// migration.Migrate()

	channels := Channel.New().All()

	for _, channel := range channels {
		chatId := channel.GetChatId()

		messages := MessageService.New().GetMessages(chatId, channel.Last_message_receive)

		for _, message := range messages.Messages {
			content := message.Content.(*client.MessageText).Text.Entities[0]

			url := content.Type.(*client.TextEntityTypeTextUrl).Url

			fmt.Println(url)
		}
	}
}
