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
	"whalego/models/Proxy"

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
	// channels := Channel.New().All()

	/* for _, channel := range channels {
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

			if !ps.isValidProxy(proxy) {
				continue;
			}

			proxyData := getProxyData(proxy)
			if proxyData == nil {
				continue
			}
		}
	} */

	channelModel := Channel.New().FindByUsername("proxymtproto")
	chatId := channelModel.GetChatId()
	messages := MessageService.New().GetMessages(chatId, channelModel.Last_message_receive)

	for _, message := range messages.Messages {
		proxy := ps.buttonMessageHandler(message)

		// proxy := "https://t.me/proxy?server=23.88.48.140&port=443&secret=DD89c92f4f14e9f5144f7f256b0feed874"

		if !ps.isValidProxy(proxy) {
			continue
		}

		proxyData, ok := ps.getProxyData(proxy)
		if ok == false {
			continue
		}

		ping, isAvailable := ps.checkProxyIsAvailable(proxyData)
		if isAvailable == false {
			continue
		}

		fmt.Println(ping, isAvailable)
	}
}

/**
* get message url from content key
*/
func (ps *ProxyService) textMessageHandler(message *client.Message) string {
	content := message.Content.(*client.MessageText).Text.Entities[0]

	url := content.Type.(*client.TextEntityTypeTextUrl).Url

	return url
}

/**
* get message url from message button key
*/
func (ps *ProxyService) buttonMessageHandler(message *client.Message) string {
	replyMarkup := message.ReplyMarkup.(*client.ReplyMarkupInlineKeyboard).Rows[0][0]

	url := replyMarkup.Type.(*client.InlineKeyboardButtonTypeUrl).Url

	return url
}

func (ps *ProxyService) isValidProxy(proxy string) bool {
	contains := false

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
func (ps *ProxyService) getProxyData(proxy string) (Proxy.Proxy, bool) {
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
		Link:   proxy,
		Server: values.Get("server"),
		Port:   int32(port),
		Secret: values.Get("secret"),
	}, true

	/* return map[string]interface{}{
		"link":   proxy,
		"server": values.Get("server"),
		"port":   int32(port),
		"secret": values.Get("secret"),
	} */
}

func (ps *ProxyService) checkProxyIsAvailable(proxy Proxy.Proxy) (string, bool) {
	// run a command to get ping of a server
	out, _ := exec.Command("ping", proxy.Server, "-c 5", "-i 3", "-w 10").Output()

	// check if server is not available
	if strings.Contains(string(out), "Destination Host Unreachable") {
		return "0", false
	}

	// get time= from result
	charindex := strings.Index(string(out), "time=")
	time := string(out[charindex+5:])
	ping := time[:4]
	pingInt, err := strconv.ParseFloat(ping, 10)

	if pingInt > 450 || err != nil {
		return "0", false
	}

	return ping, true
}
