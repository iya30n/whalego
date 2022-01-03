package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/zelenin/go-tdlib/client"
)

func main() {
	authorizer := client.ClientAuthorizer()
	go client.CliInteractor(authorizer)

	const (
		apiId   = 5339469
		apiHash = "841a53b1aabfe94d8bcf5a88a5624d6d"
	)

	authorizer.TdlibParameters <- &client.TdlibParameters{
		UseTestDc:              false,
		DatabaseDirectory:      filepath.Join(".tdlib", "database"),
		FilesDirectory:         filepath.Join(".tdlib", "files"),
		UseFileDatabase:        true,
		UseChatInfoDatabase:    true,
		UseMessageDatabase:     true,
		UseSecretChats:         false,
		ApiId:                  apiId,
		ApiHash:                apiHash,
		SystemLanguageCode:     "en",
		DeviceModel:            "Server",
		SystemVersion:          "1.0.0",
		ApplicationVersion:     "1.0.0",
		EnableStorageOptimizer: true,
		IgnoreFileNames:        false,
	}

	logVerbosity := client.WithLogVerbosity(&client.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: 0,
	})

	proxy := client.WithProxy(&client.AddProxyRequest{
		Server: "127.0.0.1",
		Port:   9050,
		Enable: true,
		Type: &client.ProxyTypeSocks5{
			Username: "",
			Password: "",
		},
	})

	tdlibClient, err := client.NewClient(authorizer, logVerbosity, proxy)
	if err != nil {
		log.Fatalf("NewClient error: %s", err)
	}

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
