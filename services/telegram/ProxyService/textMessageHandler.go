package ProxyService

import "github.com/zelenin/go-tdlib/client"

/**
* get message url from content key
 */
func textMessageHandler(message *client.Message) []string {
	contentType := message.Content.MessageContentType()
	if contentType != client.TypeMessageText && contentType != client.TypeMessage {
		return []string{}
	}

	entities := message.Content.(*client.MessageText).Text.Entities

	var proxies []string
	for _, entity := range entities {
		if entity.Type.TextEntityTypeType() != client.TypeTextEntityTypeTextUrl {
			continue
		}

		url := entity.Type.(*client.TextEntityTypeTextUrl).Url

		if url == "" || !isValidProxy(url) {
			continue
		}

		proxies = append(proxies, url)
	}

	return proxies
}
