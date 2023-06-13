package ProxyService

import (
	"os/exec"
	"strconv"
	"strings"
	"whalego/models/Proxy"
)

func checkProxyIsAvailable(proxy Proxy.Proxy) (string, bool) {
	// run a command to get ping of a server
	out, _ := exec.Command("ping", proxy.Address, "-c 5", "-i 3").Output()

	// check if server is not available
	if strings.Contains(string(out), "Destination Host Unreachable") || string(out) == "" {
		return "0", false
	}

	// get time= from result
	charindex := strings.Index(string(out), "time=")
	time := string(out[charindex+5:])
	ping := strings.TrimSpace(time[:4])
	pingInt, err := strconv.ParseFloat(ping, 32)

	if pingInt > 450 || err != nil {
		return "0", false
	}

	return ping, true
}
