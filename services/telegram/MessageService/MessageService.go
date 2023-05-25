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

func SendMarkdown(chatId int64, message string) *client.Message {
	tgConnection := connection.TdConnection(true)

	// defer connection.Close(tgConnection)

	mdMsg, err := tgConnection.ParseMarkdown(&client.ParseMarkdownRequest{
		Text: &client.FormattedText{
			Text: message,
		},
	})

	/* mdMsg, err := tgConnection.ParseTextEntities(&client.ParseTextEntitiesRequest{
		// Text: "*bold* _italic_ `code`",
		Text: message,
		ParseMode: &client.TextParseModeMarkdown{
			Version: 1,
		},
	}) */

	errorHandler.LogFile(err)

	/* tgConnection.GetBasicGroup(&client.GetBasicGroupRequest{
		BasicGroupId: chatId,
	}) */
	/*tgConnection.GetChat(&client.GetChatRequest{
		ChatId: chatId,
	})*/
	msg, err := tgConnection.SendMessage(&client.SendMessageRequest{
		ChatId: chatId,
		InputMessageContent: &client.InputMessageText{
			Text: mdMsg,
		},
	})

	errorHandler.LogFile(err)

	return msg
}

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
