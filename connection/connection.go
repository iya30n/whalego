package connection

import (
	"path/filepath"
	"whalego/errorHandler"

	"github.com/zelenin/go-tdlib/client"
)

func TdConnection(withProxy bool) *client.Client {
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

	options := []client.Option{
		logVerbosity,
	}

	if withProxy == true {
		proxy := client.WithProxy(&client.AddProxyRequest{
			/* Server: "127.0.0.1",
			Port:   9050,
			Enable: true,
			Type: &client.ProxyTypeSocks5{
				Username: "",
				Password: "",
			}, */
			Server: "trichiasis.www.Bmi.ir.Bmi--ir.ml",
			Port:   443,
			Enable: true,
			Type: &client.ProxyTypeMtproto{
				Secret: "7jK5IN_7UWQwKOL2uHjU6sFjZG4uaW50ZXJuZXQub3Jn",
			},
		})

		options = append(options, proxy)
	}

	tdlibClient, err := client.NewClient(authorizer, options...)

	errorHandler.LogFile(err)

	return tdlibClient
}
