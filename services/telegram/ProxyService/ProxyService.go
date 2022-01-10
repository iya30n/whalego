package ProxyService

import (
	"fmt"
	"whalego/models/Channel"
	"whalego/services/telegram/MessageService"

	"github.com/zelenin/go-tdlib/client"
)

type ProxyService struct {

}

func New() *ProxyService {
	return &ProxyService{}
}

func (ps *ProxyService) GetProxies() {
	channels := Channel.New().All()

	for _, channel := range channels {
		chatId := channel.GetChatId()

		messages := MessageService.New().GetMessages(chatId, channel.Last_message_receive)

		if messages.TotalCount == 0 || messages.Messages == nil{
			continue
		}

		for _, message := range messages.Messages {
			if channel.Handler == "text" {
				content := message.Content.(*client.MessageText).Text.Entities[0]

				url := content.Type.(*client.TextEntityTypeTextUrl).Url

				fmt.Println(url)
			}

			if channel.Handler == "button" {

			}
		}
	}
}