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
		// FromMessageId: fromMessage,
        Offset: 0,
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

func (cs *MessageService) SendMarkdown(chatId int64, message string) *client.Message {
	mdMsg, err := cs.tgConnection.ParseMarkdown(&client.ParseMarkdownRequest{
		Text: &client.FormattedText{
			Text: message,
		},
	})

	/* mdMsg, err := cs.tgConnection.ParseTextEntities(&client.ParseTextEntitiesRequest{
		// Text: "*bold* _italic_ `code`",
		Text: message,
		ParseMode: &client.TextParseModeMarkdown{
			Version: 1,
		},
	}) */

	errorHandler.LogFile(err)

	msg, err := cs.tgConnection.SendMessage(&client.SendMessageRequest{
        ChatId: chatId,
        InputMessageContent: &client.InputMessageText{
			Text: mdMsg,
		},
    })

	errorHandler.LogFile(err)

	return msg
}

func (cs *MessageService) DeleteMessages(chatId int64, messageIds []int64) {
	if len(messageIds) < 1 {
		return
	}

	_, err := cs.tgConnection.DeleteMessages(&client.DeleteMessagesRequest{
		ChatId: chatId,
		MessageIds: messageIds,
		Revoke: true,
	})

	errorHandler.LogFile(err)
}