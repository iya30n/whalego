package ProxyService

import (
	"whalego/Config"
	"whalego/errorHandler"
	"whalego/models/Proxy"
	"whalego/services/telegram/ChatService"
	"whalego/services/telegram/MessageService"
)

func CheckAllProxies() {
	config := Config.Get()

	chatId, err := ChatService.GetChatId(config.ChannelName)
	errorHandler.LogFile(err)

	// deleteProxies := []Proxy.Proxy{}
	var deleteProxies []uint

	deleteMessages := []int64{}

	for _, proxy := range Proxy.New().All() {
		ping, isAvailable := checkProxyIsAvailable(proxy)
		if isAvailable {
			// proxy.Ping = ping
			// proxy.Save()
			proxy.Update(map[string]interface{}{"in_channel": proxy.InChannel, "ping": ping})

			continue
		}

		if proxy.ChannelMessageId != 0 {
			deleteMessages = append(deleteMessages, proxy.ChannelMessageId)
		}

		deleteProxies = append(deleteProxies, proxy.ID)
	}

	MessageService.DeleteMessages(chatId.Id, deleteMessages)

	Proxy.New().DeleteMany(deleteProxies)
}
