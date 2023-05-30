package main

import (
	"fmt"
	"log"
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

	go tickerHandler()

	for {
		var cmd string
		fmt.Scanln(&cmd)

		switch cmd {
		case "migrate":
			migration.Migrate()
		case "test":
			fmt.Println("command getme is running...")
			con := connection.TdConnection(true)
			me, err := con.GetMe()
			errorHandler.LogFile(err)
			log.Printf("Me: %s %s [%s]", me.FirstName, me.LastName, me.Username)
		case "proxy:crawl":
			fmt.Println("command proxy:crawl is running...")
			channels := Channel.New()
			for _, channel := range channels.All() {
				ProxyService.GetProxies(&channel)
			}
		case "proxy:send":
			fmt.Println("command proxy:send is running...")
			ProxyService.SendProxy()
		case "proxy:check":
			fmt.Println("command proxy:check is running...")
			ProxyService.CheckAllProxies()
		default:
			log.Println("Invalid argument!")
		}
	}
}

func tickerHandler() {
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
}
