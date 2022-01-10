package ProxyService

import (
	"fmt"
	"strings"
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
			var proxy string

			if channel.Handler == "text" {
				proxy = ps.textMessageHandler(message)
			}

			if channel.Handler == "button" {
				proxy = ps.buttonMessageHandler(message)
			}

			if ps.isValidProxy(proxy) {
				// TODO : save
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

func (ps *ProxyService) textMessageHandler(message *client.Message) string {
	content := message.Content.(*client.MessageText).Text.Entities[0]

	url := content.Type.(*client.TextEntityTypeTextUrl).Url

	return url
}

func (ps *ProxyService) buttonMessageHandler(message *client.Message) string {
	replyMarkup := message.ReplyMarkup.(*client.ReplyMarkupInlineKeyboard).Rows[0][0]

	url := replyMarkup.Type.(*client.InlineKeyboardButtonTypeUrl).Url

	return url
}

func (ps *ProxyService) isValidProxy(proxy string) bool {
	contains := false

	for _, word := range []string{"server", "port", "proxy"} {
		contains = strings.Contains(proxy, word)
	}

	return contains
}

func (ps *ProxyService) checkAvailability(proxy string) {
	
}