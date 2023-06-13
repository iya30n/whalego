package ProxyService

import (
	"net/url"
	"strconv"
	"whalego/models/Proxy"
)

/**
* convert proxy url to Proxy model
 */
func getProxyData(proxy string) (Proxy.Proxy, bool) {
	// get query parameters from proxy url
	u, err := url.Parse(proxy)
	if err != nil {
		return Proxy.Proxy{}, false
	}

	values, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return Proxy.Proxy{}, false
	}

	// convert port to int32
	port, err := strconv.ParseInt(values.Get("port"), 10, 32)
	if err != nil {
		return Proxy.Proxy{}, false
	}

	return Proxy.Proxy{
		Url:     proxy,
		Address: values.Get("server"),
		Port:    int32(port),
		Secret:  values.Get("secret"),
	}, true

	/* return map[string]interface{}{
		"url":   proxy,
		"address": values.Get("server"),
		"port":   int32(port),
		"secret": values.Get("secret"),
	} */
}
