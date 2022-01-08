package MessageService

import (
	"whalego/errorHandler"

	"whalego/connection"

	"github.com/zelenin/go-tdlib/client"
)

type MessageService struct {
	tgConnection *client.Client
}

func New() *MessageService {
	return &MessageService{
		tgConnection: connection.TdConnection(true),
	}
}

func (ms *MessageService) GetMessages(chatId int64, fromMessage int64) *client.Messages {
	result, err := ms.tgConnection.GetChatHistory(&client.GetChatHistoryRequest{
        ChatId: chatId,
		FromMessageId: fromMessage,
        // Offset: 0,
        Limit: 99,
        OnlyLocal: false,
    })

	errorHandler.LogFile(err)

	return result
}

func (cs *MessageService) SendMessage(chatId int64, message client.InputMessageContent) *client.Message {
	msg, err := cs.tgConnection.SendMessage(&client.SendMessageRequest{
        ChatId: chatId,
        InputMessageContent: message,
    })

	errorHandler.LogFile(err)

	return msg
}