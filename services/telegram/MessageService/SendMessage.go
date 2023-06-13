package MessageService

import (
	"whalego/errorHandler"

	"whalego/connection"

	"github.com/zelenin/go-tdlib/client"
)

func SendMessage(chatId int64, message client.InputMessageContent) *client.Message {
	tgConnection := connection.TdConnection(true)

	// defer connection.Close(tgConnection)

	msg, err := tgConnection.SendMessage(&client.SendMessageRequest{
		ChatId:              chatId,
		InputMessageContent: message,
	})

	errorHandler.LogFile(err)

	return msg
}