package ProxyService

import (
	"fmt"
	"strings"
	"whalego/Config"
	"whalego/errorHandler"
	"whalego/models/Proxy"
	"whalego/services/telegram/ChatService"
	"whalego/services/telegram/MessageService"
)

func SendProxy() {
	config := Config.Get()
	chatId, err := ChatService.GetChatId(config.ChannelName)

	errorHandler.LogFile(err)

	var availableProxy Proxy.Proxy

	for _, p := range Proxy.New().GetNotInChannel(5) {
		if _, ok := checkProxyIsAvailable(p); ok {
			availableProxy = p
			break
		}
	}

	if len(availableProxy.Url) == 0 {
		return
	}

	proxyUrl := strings.Replace(availableProxy.Url, " ", "%20", -1)

	proxyMessage := "server: " + availableProxy.Address + "\nport: %d\nping: **" + availableProxy.Ping + "**\n\n ▶️[ Connect ](" + proxyUrl + ")◀️\n➖➖➖➖➖➖➖➖➖➖\n🔽**پروکسی های بیشتر**🔽\n🆔 @whaleproxies"
	proxyMessage = fmt.Sprintf(proxyMessage, availableProxy.Port)

	// proxyMessage := "server: %d \nport: %d\nping: **%d**\n\n [▶️   Connect   ◀️](%d) \n➖➖➖➖➖➖➖➖➖➖\n🔽**پروکسی های بیشتر**🔽\n🆔 @whaleproxies"
	// proxyMessage := "server: `%d` \nport: `%d`\nping: **`%d`**\n\n [▶️   Connect   ◀️](`%d`) \n➖➖➖➖➖➖➖➖➖➖\n🔽**پروکسی های بیشتر**🔽\n🆔 @whaleproxies"
	// proxyMessage = fmt.Sprint(proxyMessage, availableProxy.Address, availableProxy.Port, availableProxy.Ping, availableProxy.Url)

	sentMessage := MessageService.SendMarkdown(chatId.Id, proxyMessage)

	availableProxy.Update(map[string]interface{}{
		"in_channel":         true,
		"channel_message_id": sentMessage.Id,
	})
}
