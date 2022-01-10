package main

import "whalego/services/telegram/ProxyService"

// "whalego/database/migration"

func main() {
	// migration.Migrate()

	ProxyService.New().GetProxies()
}
