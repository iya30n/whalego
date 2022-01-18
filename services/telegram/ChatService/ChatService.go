package ChatService

import (
	"whalego/connection"

	"github.com/zelenin/go-tdlib/client"
)

type ChatService struct {
	tgConnection *client.Client
}

func New() *ChatService {
	return &ChatService{
		tgConnection: connection.TdConnection(true),
	}
}

func (cs *ChatService) GetChatId(username string) (*client.Chat, error) {
	cs.tgConnection.GetChats(&client.GetChatsRequest{})
	return cs.tgConnection.SearchPublicChat(&client.SearchPublicChatRequest{
		Username: username,
	})
}