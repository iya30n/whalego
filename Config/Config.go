package Config

import (
	"encoding/json"
	"whalego/file"
)

// develop
/* var Database map[string]string = map[string]string{
	"HOST":     "127.0.0.1",
	"PORT":     "3306",
	"DBNAME":   "whaleproxy",
	"USERNAME": "root",
	"PASS":     "",
}

var Telegram map[string]string = map[string]string{
	"whale_channel": "whale_test",
} */

// production
/* var Database map[string]string = map[string]string{
	"HOST":     "127.0.0.1",
	"PORT":     "3306",
	"DBNAME":   "whaleproxy",
	"USERNAME": "whaleproxy",
	"PASS":     "NFkui2oct4",
}

var Telegram map[string]string = map[string]string{
	"whale_channel":     "whaleproxies",
} */


type api struct {
	ApiId   int32    `json:"api_id"`
	ApiHash string `json:"api_hash"`
}

type database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"db_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type proxy struct {
	Server string `json:"server"`
	Port   int32 `json:"port"`
	Secret string `json:"secret"`
}

type Config struct {
	ChannelName string `json:"channel_name"`
	Api      api      `json:"api"`
	Database database `json:"database"`
	Proxy    proxy    `json:"proxy"`
}

func Get() Config {
	var config Config
	json.Unmarshal(file.Read("./config.json"), &config)

	return config
}