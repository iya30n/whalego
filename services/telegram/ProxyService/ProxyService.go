package ProxyService

import (
	// "errors"
	"fmt"
	"net/url"
	"strings"
	"whalego/connection"
	"whalego/errorHandler"
	// "whalego/models/Channel"
	// "whalego/services/telegram/MessageService"

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

	/* channelModel := Channel.New().FindByUsername("proxymtproto")
	chatId := channelModel.GetChatId()
	messages := MessageService.New().GetMessages(chatId, channelModel.Last_message_receive) */

	// for _, message := range messages.Messages {
		// proxy := ps.buttonMessageHandler(message)

		proxy := "https://t.me/proxy?server=23.88.48.140&port=443&secret=DD89c92f4f14e9f5144f7f256b0feed874"

		/* if !ps.isValidProxy(proxy) {
			continue
		} */	

		proxyData := ps.getProxyData(proxy)
		/* if proxyData == nil {
			continue
		} */

		ps.checkProxyIsAvailable(proxyData)
	// }
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

	for _, word := range []string{"proxy", "server", "port"} {
		contains = strings.Contains(proxy, word)

		if contains == false {
			return false
		}
	}

	return contains
}

func (ps *ProxyService) getProxyData(proxy string) map[string]interface{} {
	u, err := url.Parse(proxy)
	errorHandler.LogFile(err)

	values, err := url.ParseQuery(u.RawQuery)
	errorHandler.LogFile(err)

	return map[string]interface{}{
		"link": proxy,
		"server": values.Get("server"),
		"port": values.Get("port"),
		"secret": values.Get("secret"),
	}

	/* data, err := url.ParseQuery(proxy)
	errorHandler.LogFile(err)

	server, ok := data["https://t.me/proxy?server"]
	if !ok {
		return nil
	}

	port, ok := data["port"]
	if !ok {
		return nil
	}

	secret, ok := data["secret"]
	if !ok {
		return nil
	}

	return map[string]interface{}{
		"link": proxy,
		"server": server[0],
		"port": port[0],
		"secret": secret[0],
	} */
}

func (ps *ProxyService) checkProxyIsAvailable(proxy map[string]interface{}) {
	ok,err := connection.TdConnection(true).TestProxy(&client.TestProxyRequest{
		// Server: fmt.Sprint(proxy["server"]),
		Server: "23.88.48.140",
		Port: 443,
		Type: &client.ProxyTypeMtproto{
			Secret: "DD89c92f4f14e9f5144f7f256b0feed874",
		},
		Timeout: 12000,
	})

	/* ok,err := connection.TdConnection(true).TestProxy(&client.TestProxyRequest{
		// Server: fmt.Sprint(proxy["server"]),
		Server: proxy["server"].(string),
		Port: proxy["port"].(int32),
		Type: &client.ProxyTypeMtproto{},
	}) */

	// connection.TdConnection(false).PingProxy(&client.PingProxyRequest{})

	errorHandler.LogFile(err)

	fmt.Println(ok)
}