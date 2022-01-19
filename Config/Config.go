package Config

// develop
/* var Database map[string]string = map[string]string{
	"HOST":     "127.0.0.1",
	"PORT":     "3306",
	"DBNAME":   "whaleproxy",
	"USERNAME": "root",
	"PASS":     "",
}

var Telegram map[string]string = map[string]string{
	"whale_channel":     "whale_test",
} */


// production
var Database map[string]string = map[string]string{
	"HOST":     "127.0.0.1",
	"PORT":     "3306",
	"DBNAME":   "whaleproxy",
	"USERNAME": "whaleproxy",
	"PASS":     "NFkui2oct4",
}

var Telegram map[string]string = map[string]string{
	"whale_channel":     "whaleproxies",
}