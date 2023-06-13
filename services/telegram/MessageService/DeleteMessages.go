package MessageService

import (
	"whalego/errorHandler"

	"whalego/connection"

	"github.com/zelenin/go-tdlib/client"
)

func DeleteMessages(chatId int64, messageIds []int64) {
	tgConnection := connection.TdConnection(true)

	// defer connection.Close(tgConnection)

	if len(messageIds) < 1 {
		return
	}

	_, err := tgConnection.DeleteMessages(&client.DeleteMessagesRequest{
		ChatId:     chatId,
		MessageIds: messageIds,
		Revoke:     true,
	})

	errorHandler.LogFile(err)
}
