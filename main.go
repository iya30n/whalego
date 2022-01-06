package main

import (
	"log"
	"os"
	"whalego/models/Channel"
	// "whalego/connection"
	// "whalego/services/telegram/ChatService"
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
    /* tdlibClient := connection.TdConnection(true)

    chat := ChatService.New(tdlibClient)

    messages := chat.GetMessages("be3t_proxy")

    mjson, err := messages.MarshalJSON()
    if err != nil {
        panic(err)
    }

    newFile(mjson) */

    channel := Channel.New().FindByUsername("@darkproxyy")

    log.Println(channel.Handler)
}
