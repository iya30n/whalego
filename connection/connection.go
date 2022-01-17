package connection

import (
	"path/filepath"
	"whalego/errorHandler"

	"github.com/zelenin/go-tdlib/client"
)

var singletonConnection *client.Client
var isset bool = false

func init() {
	singletonConnection = TdConnection(true)
	isset = true
}

func TdConnection(withProxy bool) *client.Client {
	if isset {
		return singletonConnection
	} else {
		return makeConnection(withProxy)
	}
}

func makeConnection(withProxy bool) *client.Client {
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
		UseChatInfoDatabase:    false,
		UseMessageDatabase:     false,
		UseSecretChats:         false,
		ApiId:                  apiId,
		ApiHash:                apiHash,
		SystemLanguageCode:     "en",
		DeviceModel:            "Server",
		SystemVersion:          "1.0.0",
		ApplicationVersion:     "1.0.0",
		EnableStorageOptimizer: false,
		IgnoreFileNames:        false,
	}

	logVerbosity := client.WithLogVerbosity(&client.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: 0,
	})

	options := []client.Option{
		logVerbosity,
	}

	if withProxy {
		proxy := client.WithProxy(&client.AddProxyRequest{
			Server: "127.0.0.1",
			Port:   9050,
			Enable: true,
			Type: &client.ProxyTypeSocks5{
				Username: "",
				Password: "",
			},
			/* Server: "www.cloudflare.tattoo",
			Port:   443,
			Enable: true,
			Type: &client.ProxyTypeMtproto{
				Secret: "dd00000000000000000000000000000000",
			}, */
		})

		options = append(options, proxy)
	}

	tdlibClient, err := client.NewClient(authorizer, options...)

	errorHandler.LogFile(err)

	return tdlibClient
}

func Close(connection *client.Client) {
	isset = false

	connection.Close()
}
