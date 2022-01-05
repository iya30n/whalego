package Channel

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model
	username             string `gorm: "NOT NULL;size:256"`
	ChatId               int64  `gorm: "int, unique;"`
	last_message_receive int    `gorm: "int;"`
	handler              string `gorm: "NOT NULL;size:30"`
}

func New() *Channel {
	return &Channel{}
}

func (c *Channel) All() {
	dsn := "root:@tcp(127.0.0.1:3306)/whaleproxy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(c)

	// TODO: use err handler here
	if err != nil {
		panic(err)
	}

	// channel := &Channel{}
	// result, _ := db.Select("username", "handler").Find(&channel).Rows()

	// log.Printf("%+v", c)
	/* for result.Next() {

		// var channel Channel

		err := result.Scan(
			&channel.username,
			&channel.ChatId,
			&channel.last_message_receive,
			&channel.handler,
			&channel.ID,
			&channel.CreatedAt,
			&channel.UpdatedAt,
			&channel.DeletedAt,
		)

		if err != nil {
			panic(err)
		}

		log.Println()
	} */

	// log.Printf("%+v", result)

	var channels []Channel

	db.Find(&channels)
	
	for i:=0; i < len(channels); i++ {
		log.Println(channels[i].username)
	}
}
