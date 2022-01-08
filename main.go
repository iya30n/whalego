package main

import (
	"encoding/json"
	// "fmt"
	"os"

	// "whalego/errorHandler"
	"whalego/errorHandler"
	"whalego/models/Channel"
	"whalego/services/telegram/MessageService"
	// "whalego/database/migration"
)

func newFile (val []byte) {
    f, err := os.Create("./result.json")
    if err != nil {
        panic(err)
    }

    defer f.Close()

    f.Write(val)

    f.Sync()
}

func main() {
    // migration.Migrate()

    channels := Channel.New().All()

    for _, channel := range channels {
        chatId := channel.GetChatId()

        messages := MessageService.New().GetMessages(chatId, channel.Last_message_receive)

       /*  jsonMessages, err := messages.MarshalJSON()

        errorHandler.LogFile(err)

        newFile(jsonMessages) */

        marshalMsg, err := messages.MarshalJSON()

        errorHandler.LogFile(err)

        var unmarshalMsg map[string]interface{}

        json.Unmarshal(marshalMsg, &unmarshalMsg)

        print(unmarshalMsg["@type"])

        /* for _, message := range marshalMsg {
            fmt.Println(message)
        } */

        /* for _, message := range messages.Messages {
            fmt.Println(message.Content)
        }*/
    }
}
