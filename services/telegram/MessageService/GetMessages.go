package MessageService

import (
	"whalego/errorHandler"

	"whalego/connection"

	"github.com/zelenin/go-tdlib/client"
)

func GetMessages(chatId int64, fromMessage int64) *client.Messages {
	tgConnection := connection.TdConnection(true)
	// defer connection.Close(tgConnection)

	/* tgConnection.GetMessages(&client.GetMessagesRequest{

	}) */

	result, err := tgConnection.GetChatHistory(&client.GetChatHistoryRequest{
		ChatId: chatId,
		// FromMessageId: fromMessage,
		// Offset:    0,
		FromMessageId: 0,
		Limit:     99,
		OnlyLocal: false,
	})

	errorHandler.LogFile(err)

	return result
}