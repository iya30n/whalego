package ChatService

import (
	"os"
	"whalego/errorHandler"
	"whalego/models/Channel"

	"github.com/zelenin/go-tdlib/client"
)

type ChatService struct {
	tgConnection *client.Client
}

func New(connection *client.Client) *ChatService {
	return &ChatService{
		tgConnection: connection,
	}
}

func (cs *ChatService) GetChatId(username string) int64 {
	channel := Channel.New().FindByUsername(username)

	if channel.ChatId != 0 {
		return channel.ChatId
	}

	chat, err := cs.tgConnection.SearchPublicChat(&client.SearchPublicChatRequest{
		Username: username,
	})

	if err.Error() == "USERNAME_NOT_OCCUPIED" && chat == nil {
		channel.Delete()
		os.Exit(1)
	}

	errorHandler.LogFile(err)

	channel.Update(map[string]interface{} {
		"chat_id": chat.Id,
	})

	return chat.Id
}