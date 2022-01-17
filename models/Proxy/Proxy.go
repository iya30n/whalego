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

func (p *Proxy) All() []Proxy{
	db := database.Connect()

	defer database.Close(db)

	var proxies []Proxy

	db.Find(&proxies)

	return proxies
}

func (p *Proxy) GetLimit(limit int) []Proxy{
	db := database.Connect()

	defer database.Close(db)

	var proxies []Proxy

	db.Limit(limit).Find(&proxies)

	return proxies
}

func (p *Proxy) GetNotInChannel(limit int) []Proxy{
	db := database.Connect()

	defer database.Close(db)

	var proxies []Proxy

	db.Where("in_channel = ?", false).Limit(limit).Find(&proxies)

	return proxies
}

func (p *Proxy) Create(data map[string]interface{}) {
	db := database.Connect()

	defer database.Close(db)

	if _, ok := data["in_channel"]; ok == false {
		data["in_channel"] = false
	}

	db.Model(&p).Create(data)
}

func (p *Proxy) Update(data map[string]interface{}) {
	db := database.Connect()

	defer database.Close(db)

	if _, ok := data["in_channel"]; ok == false {
		data["in_channel"] = false
	}

	db.Model(&p).UpdateColumns(data)
}

func (p *Proxy) Save() {
	db := database.Connect()

	defer database.Close(db)

	p.InChannel = false

	db.Create(&p)
}

func (p *Proxy) Exists() bool {
	db := database.Connect()

	defer database.Close(db)

	var proxies []Proxy
	db.Where("address = ?", p.Address).Find(&proxies)

	return len(proxies) > 0
}

func (p *Proxy) Delete() {
	db := database.Connect()

	defer database.Close(db)

	db.Unscoped().Delete(&p)
}