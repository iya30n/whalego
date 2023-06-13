package ProxyService

import "github.com/zelenin/go-tdlib/client"

/**
* get message url from message button key
 */
func buttonMessageHandler(message *client.Message) []string {
	var proxies []string

	if message.ReplyMarkup == nil {
		return proxies
	}

	rows := message.ReplyMarkup.(*client.ReplyMarkupInlineKeyboard).Rows
	
	for _, row := range rows {
		for _, btn := range row {
			if btn.Type.InlineKeyboardButtonTypeType() != client.TypeInlineKeyboardButtonTypeUrl {
				continue
			}
		
			url := btn.Type.(*client.InlineKeyboardButtonTypeUrl).Url
	
			if url == "" || !isValidProxy(url) {
				continue
			}
	
			proxies = append(proxies, url)
		}
	}

	return proxies
}
