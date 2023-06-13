package ProxyService

import (
	"fmt"
	"whalego/models/Channel"
	"whalego/services/telegram/MessageService"
)

func GetProxies(channel *Channel.Channel) {
	fmt.Println("checking channel " + channel.Username)

	chatId := channel.GetChatId()
	if chatId == 0 {
		return
	}

	messages := MessageService.GetMessages(chatId, channel.Last_message_receive)

	if messages.Messages == nil {
		return
	}

	var proxies []string

	for _, message := range messages.Messages {
		if channel.Handler == "text" {
			proxies = append(proxies, textMessageHandler(message)...)
		}

		if channel.Handler == "button" {
			proxies = append(proxies, buttonMessageHandler(message)...)
		}
	}

	for _, proxy := range proxies {
		proxyData, ok := getProxyData(proxy)
		if !ok {
			continue
		}

		if proxyData.Exists() {
			continue
		}

		ping, isAvailable := checkProxyIsAvailable(proxyData)
		if !isAvailable {
			continue
		}

		proxyData.Ping = ping

		proxyData.Save()

		fmt.Println("one proxy saved")
	}
}
