package ChatService

import (
	"whalego/connection"

	"github.com/zelenin/go-tdlib/client"
)

func GetChatId(username string) (*client.Chat, error) {
	tgConnection := connection.TdConnection(true)

	defer connection.Close(tgConnection)
	// tgConnection.GetChats(&client.GetChatsRequest{})
	return tgConnection.SearchPublicChat(&client.SearchPublicChatRequest{
		Username: username,
	})
}
