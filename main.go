package main

import (
	"fmt"
	"os"
	"whalego/database/migration"
	"whalego/models/Channel"
	"whalego/services/telegram/ProxyService"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("what can i do for you?")
		os.Exit(1)
	}

	if os.Args[1] == "migrate" {
		migration.Migrate()
		os.Exit(1)
	}

	if os.Args[1] == "proxy:crawler" {
		channels := Channel.New()
		for _, channel := range channels.All() {
			ProxyService.GetProxies(&channel)
		}

		os.Exit(1)
	}

	if os.Args[1] == "proxy:send" {
		ProxyService.SendProxy()
		os.Exit(1)
	}

	if os.Args[1] == "proxy:check" {
		ProxyService.CheckAllProxies()
		os.Exit(1)
	}
}
