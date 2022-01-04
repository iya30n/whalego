package main

import (
	"fmt"
	"log"
	"whalego/connection"

	"github.com/zelenin/go-tdlib/client"
)

func main() {
    tdlibClient := connection.TdConnection(true)

    user, err := tdlibClient.SearchPublicChat(&client.SearchPublicChatRequest{
        Username: "iya30n",
    })

    if err != nil {
		log.Fatalf("Get user error: %s", err)
    }

    msg, err := tdlibClient.SendMessage(&client.SendMessageRequest{
        ChatId: user.Id,
        InputMessageContent: &client.InputMessageText{
            Text: &client.FormattedText{
                Text: "salam from whalego",
            },
        },
    })

    if err != nil {
		log.Fatalf("Get user error: %s", err)
    }

    fmt.Println(msg)
}
