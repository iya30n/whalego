package ChatService

import (
	"whalego/errorHandler"

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

	// TODO: check if chat not found

	errorHandler.LogFile(err)
	}

	channel.Update(map[string]interface{} {
		"chat_id": chat.Id,
	})

	return chat.Id
}

func (cs *ChatService) GetMessages(username string) *client.Messages {
	chatId := cs.GetChatId(username)

	result, err := cs.tgConnection.GetChatHistory(&client.GetChatHistoryRequest{
        ChatId: chatId,
        Offset: 0,
        Limit: 99,
        OnlyLocal: false,
    })

	errorHandler.LogFile(err)

	return result
}

func (cs *ChatService) SendMessage(username string, message client.InputMessageContent) *client.Message {
	chatId := cs.GetChatId(username)

	msg, err := cs.tgConnection.SendMessage(&client.SendMessageRequest{
        ChatId: chatId,
        InputMessageContent: message,
    })

	errorHandler.LogFile(err)

	return msg
}