package Proxy

import (
	"whalego/database"

	"gorm.io/gorm"
)

type Proxy struct {
	gorm.Model
	Url              string
	Address          string `gorm: "NOT NULL;"`
	Port             int32  `gorm: "NOT NULL;"`
	Secret           string
	Ping             string
	InChannel        bool `gorm: "NOT NULL;"`
	ChannelMessageId int64
}

func New() *Proxy {
	return &Proxy{}
}

func (p *Proxy) Create(data map[string]interface{}) {
	db := database.Connect()

	if _, ok := data["in_channel"]; ok == false {
		data["in_channel"] = false
	}

	db.Model(&p).Create(data)
}

func (p *Proxy) Save() {
	db := database.Connect()

	p.InChannel = false

	db.Create(&p)
}

func (p *Proxy) Exists() bool {
	db := database.Connect()

	var proxies []Proxy
	db.Where("address = ?", p.Address).Find(&proxies)

	return len(proxies) > 0
}