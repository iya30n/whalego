package main

import (
	"fmt"
	"log"
	"whalego/connection"
	"whalego/database/migration"
	"whalego/errorHandler"
	"whalego/models/Channel"
	"whalego/services/telegram/ChatService"
	"whalego/services/telegram/MessageService"
	"whalego/services/telegram/ProxyService"
)

func main() {
	for {
		var command string

		fmt.Scanln(&command)

		if len(command) == 0 {
			break
		}

		if command == "migrate" {
			fmt.Println("command migrate is running...")
			migration.Migrate()
			continue
		}

		if command == "proxy:crawler" {
			fmt.Println("command proxy:crawler is running...")
			channels := Channel.New()
			for _, channel := range channels.All() {
				ProxyService.GetProxies(&channel)
			}

			continue
		}

		if command == "proxy:send" {
			fmt.Println("command proxy:send is running...")
			ProxyService.SendProxy()
			continue
		}

		if command == "proxy:check" {
			fmt.Println("command proxy:check is running...")
			ProxyService.CheckAllProxies()
			continue
		}

		if command == "test" {
			fmt.Println("command test is running...")
			chatID, err := ChatService.GetChatId("iya30n")
			errorHandler.LogFile(err)
			MessageService.SendMarkdown(chatID.Id, "salam from whale")

			continue
		}

		if command == "getme" {
			fmt.Println("command getme is running...")
			con := connection.TdConnection(true)
			me, err := con.GetMe()
			errorHandler.LogFile(err)
			log.Printf("Me: %s %s [%s]", me.FirstName, me.LastName, me.Username)
			continue
		}
	}
}
