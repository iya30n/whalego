package ProxyService

import "strings"

func isValidProxy(proxy string) bool {
	contains := false

	if strings.Contains(proxy, "MISSING") || strings.Contains(proxy, "(") {
		return false
	}

	for _, word := range []string{"proxy", "server", "port"} {
		contains = strings.Contains(proxy, word)

		if contains == false {
			return false
		}
	}

	return contains
}
