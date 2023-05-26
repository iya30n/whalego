package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"whalego/connection"
	"whalego/database/migration"
	"whalego/errorHandler"
	"whalego/models/Channel"
	// "whalego/services/telegram/ChatService"
	// "whalego/services/telegram/MessageService"
	"whalego/services/telegram/ProxyService"
)

func main() {

	if len(os.Args) > 1 {
		fmt.Println("command migrate is running...")

		switch os.Args[1] {
		case "migrate":
			migration.Migrate()
		case "test":
			fmt.Println("command getme is running...")
			con := connection.TdConnection(true)
			me, err := con.GetMe()
			errorHandler.LogFile(err)
			log.Printf("Me: %s %s [%s]", me.FirstName, me.LastName, me.Username)
		default:
			log.Println("Invalid argument!")
		}

		return
	}

	crawlerTicker := time.NewTicker(time.Hour * 2)

	sendTicker := time.NewTicker(time.Hour)

	checkTicker := time.NewTicker(time.Hour * 6)

	for {
		select {
		case <-crawlerTicker.C:
			channels := Channel.New()
			for _, channel := range channels.All() {
				ProxyService.GetProxies(&channel)
			}

		case <-sendTicker.C:
			ProxyService.SendProxy()

		case <-checkTicker.C:
			ProxyService.CheckAllProxies()
		}
	}

	/* fmt.Println("command test is running...")
	chatID, err := ChatService.GetChatId("iya30n")
	errorHandler.LogFile(err)
	MessageService.SendMarkdown(chatID.Id, "salam from whale")

	fmt.Println("command getme is running...")
	con := connection.TdConnection(true)
	me, err := con.GetMe()
	errorHandler.LogFile(err)
	log.Printf("Me: %s %s [%s]", me.FirstName, me.LastName, me.Username) */

}
