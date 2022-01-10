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
				ps.textMessageHandler(message)
			}

			if channel.Handler == "button" {
				ps.buttonMessageHandler(message)
			}
		}
	}

	/* channelModel := Channel.New().FindByUsername("proxymtproto")
	chatId := channelModel.GetChatId()
	messages := MessageService.New().GetMessages(chatId, channelModel.Last_message_receive)

	for _, message := range messages.Messages {
		ps.buttonMessageHandler(message)
	} */

}

func (ps *ProxyService) textMessageHandler(message *client.Message) {
	content := message.Content.(*client.MessageText).Text.Entities[0]

	url := content.Type.(*client.TextEntityTypeTextUrl).Url

	fmt.Println(url)
}

func (ps *ProxyService) buttonMessageHandler(message *client.Message) {
	replyMarkup := message.ReplyMarkup.(*client.ReplyMarkupInlineKeyboard).Rows[0][0]

	url := replyMarkup.Type.(*client.InlineKeyboardButtonTypeUrl).Url

	fmt.Println(url)
}
