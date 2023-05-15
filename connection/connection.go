package connection

import (
	"log"
	"path/filepath"
	"sync"
	"whalego/Config"
	"whalego/errorHandler"

	"github.com/zelenin/go-tdlib/client"
)

var doOnce sync.Once
var singletonConnection *client.Client

func TdConnection(withProxy bool) *client.Client {
	doOnce.Do(func() {
		singletonConnection = makeConnection(withProxy)
	})

	return singletonConnection
}

func makeConnection(withProxy bool) *client.Client {
	config := Config.Get()

	authorizer := client.ClientAuthorizer()
	go client.CliInteractor(authorizer)

	authorizer.TdlibParameters <- &client.TdlibParameters{
		UseTestDc:              false,
		DatabaseDirectory:      filepath.Join(".tdlib", "database"),
		FilesDirectory:         filepath.Join(".tdlib", "files"),
		UseFileDatabase:        true,
		UseChatInfoDatabase:    true,
		UseMessageDatabase:     true,
		UseSecretChats:         false,
		ApiId:                  config.Api.ApiId,
		ApiHash:                config.Api.ApiHash,
		SystemLanguageCode:     "en",
		DeviceModel:            "Server",
		SystemVersion:          "1.0.0",
		ApplicationVersion:     "1.0.0",
		EnableStorageOptimizer: true,
		IgnoreFileNames:        false,
	}

	_, err := client.SetLogVerbosityLevel(&client.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: 1,
	})
	errorHandler.LogFile(err)

	var proxyType client.ProxyType
	if len(config.Proxy.Secret) == 0 {
		proxyType = &client.ProxyTypeSocks5{
			Username: "",
			Password: "",
		}
	} else {
		proxyType = &client.ProxyTypeMtproto{
			Secret: config.Proxy.Secret,
		}
	}

	// if withProxy {
	proxy := client.WithProxy(&client.AddProxyRequest{
		Server: config.Proxy.Server,
		Port:   config.Proxy.Port,
		Enable: true,
		Type: proxyType,
	})

	tdlibClient, err := client.NewClient(authorizer, proxy)
	errorHandler.LogFile(err)

	optionValue, err := tdlibClient.GetOption(&client.GetOptionRequest{
		Name: "version",
	})
	errorHandler.LogFile(err)

	log.Printf("TDLib version: %s", optionValue.(*client.OptionValueString).Value)

	return tdlibClient
}

func Close(connection *client.Client) {
	// connection.Stop()
	// connection.Close()
}
