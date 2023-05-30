package ProxyService

import (
	// "errors"
	"fmt"
	"net/url"
	"os/exec"
	"strconv"
	"strings"

	// "whalego/connection"
	// "whalego/errorHandler"
	"whalego/Config"
	"whalego/errorHandler"
	"whalego/models/Proxy"

	"whalego/models/Channel"
	"whalego/services/telegram/ChatService"
	"whalego/services/telegram/MessageService"

	"github.com/zelenin/go-tdlib/client"
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

	proxyMessage := "server: " + availableProxy.Address + "\nport: %d\nping: **" + availableProxy.Ping + "**\n\n ▶️[ Connect ](" + availableProxy.Url + ")◀️\n➖➖➖➖➖➖➖➖➖➖\n🔽**پروکسی های بیشتر**🔽\n🆔 @whaleproxies"
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

func CheckAllProxies() {
	config := Config.Get()

	chatId, err := ChatService.GetChatId(config.ChannelName)
	errorHandler.LogFile(err)

	deleteProxies := []Proxy.Proxy{}

	deleteMessages := []int64{}

	for _, proxy := range Proxy.New().All() {
		ping, isAvailable := checkProxyIsAvailable(proxy)
		if isAvailable {
			proxy.Ping = ping
			proxy.Save()

			continue
		}

		if proxy.ChannelMessageId != 0 {
			deleteMessages = append(deleteMessages, proxy.ChannelMessageId)
		}

		deleteProxies = append(deleteProxies, proxy)
	}

	MessageService.DeleteMessages(chatId.Id, deleteMessages)

	Proxy.New().DeleteMany(deleteProxies)
}

/**
* get message url from content key
 */
func textMessageHandler(message *client.Message) string {
	contentType := message.Content.MessageContentType()
	if contentType != client.TypeMessageText && contentType != client.TypeMessage {
		return ""
	}

	entities := message.Content.(*client.MessageText).Text.Entities

	var url string
	for _, entity := range entities {
		if entity.Type.TextEntityTypeType() != client.TypeTextEntityTypeTextUrl {
			continue
		}

		url = entity.Type.(*client.TextEntityTypeTextUrl).Url

		if url != "" {
			break
		}
	}

	return url
}

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

func isValidProxy(proxy string) bool {
	contains := false

	if strings.Contains(proxy, "MISSING") || strings.Contains(proxy, "(") {
		return false
	}

	for _, word := range []string{"proxy", "server", "port"} {
		contains = strings.Contains(proxy, word)

		if contains == false {
			return false
		}
	}

	return contains
}

/**
* convert proxy url to Proxy model
 */
func getProxyData(proxy string) (Proxy.Proxy, bool) {
	// get query parameters from proxy url
	u, err := url.Parse(proxy)
	if err != nil {
		return Proxy.Proxy{}, false
	}

	values, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return Proxy.Proxy{}, false
	}

	// convert port to int32
	port, err := strconv.ParseInt(values.Get("port"), 10, 32)
	if err != nil {
		return Proxy.Proxy{}, false
	}

	return Proxy.Proxy{
		Url:     proxy,
		Address: values.Get("server"),
		Port:    int32(port),
		Secret:  values.Get("secret"),
	}, true

	/* return map[string]interface{}{
		"url":   proxy,
		"address": values.Get("server"),
		"port":   int32(port),
		"secret": values.Get("secret"),
	} */
}

func checkProxyIsAvailable(proxy Proxy.Proxy) (string, bool) {
	// run a command to get ping of a server
	out, _ := exec.Command("ping", proxy.Address, "-c 5", "-i 3").Output()

	// check if server is not available
	if strings.Contains(string(out), "Destination Host Unreachable") || string(out) == "" {
		return "0", false
	}

	// get time= from result
	charindex := strings.Index(string(out), "time=")
	time := string(out[charindex+5:])
	ping := strings.TrimSpace(time[:4])
	pingInt, err := strconv.ParseFloat(ping, 32)

	if pingInt > 450 || err != nil {
		return "0", false
	}

	return ping, true
}
