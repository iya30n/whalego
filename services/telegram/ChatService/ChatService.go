package ChatService

import (
	"log"

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
	// TODO: return chat_id from db->channels; else get it from telegram and update the channel is null chat_id
	chat, err := cs.tgConnection.SearchPublicChat(&client.SearchPublicChatRequest{
		Username: username,
	})

	// TODO: check if chat not found

	// TODO: chat error handler here.
	if err != nil {
		panic(err)
	}

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

	if err != nil {
		log.Fatalf("Get user error: %s", err)
    }

	return result
}

func (cs *ChatService) SendMessage(username string, message client.InputMessageContent) *client.Message {
	chatId := cs.GetChatId(username)

	msg, err := cs.tgConnection.SendMessage(&client.SendMessageRequest{
        ChatId: chatId,
        InputMessageContent: message,
    })

	// TODO: use error handler here
	if err != nil {
		panic(err)
	}

	return msg
}