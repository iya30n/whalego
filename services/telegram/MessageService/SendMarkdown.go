package MessageService

import (
	"whalego/errorHandler"

	"whalego/connection"

	"github.com/zelenin/go-tdlib/client"
)

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