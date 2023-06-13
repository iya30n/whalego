package ProxyService

import "github.com/zelenin/go-tdlib/client"

/**
* get message url from message button key
 */
func buttonMessageHandler(message *client.Message) string {

	if message.ReplyMarkup == nil {
		return ""
	}

	replyMarkup := message.ReplyMarkup.(*client.ReplyMarkupInlineKeyboard).Rows[0][0]

	if replyMarkup.Type.InlineKeyboardButtonTypeType() != client.TypeInlineKeyboardButtonTypeUrl {
		return ""
	}

	url := replyMarkup.Type.(*client.InlineKeyboardButtonTypeUrl).Url

	return url
}
