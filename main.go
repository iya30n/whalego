package main

import (
	"whalego/database/migration"
	"whalego/services/telegram/ProxyService"
)

func main() {
	migration.Migrate()

	ProxyService.New().GetProxies()
}
