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

	if messages.TotalCount == 0 || messages.Messages == nil {
		return
	}

	for _, message := range messages.Messages {
		var proxy string

		if channel.Handler == "text" {
			proxy = textMessageHandler(message)
		}

		if channel.Handler == "button" {
			proxy = buttonMessageHandler(message)
		}

		if !isValidProxy(proxy) {
			continue
		}

		proxyData, ok := getProxyData(proxy)
		if ok == false {
			continue
		}

		if proxyData.Exists() {
			continue
		}

		ping, isAvailable := checkProxyIsAvailable(proxyData)
		if isAvailable == false {
			continue
		}

		proxyData.Ping = ping

		proxyData.Save()

		fmt.Println("one proxy saved")
	}
}
