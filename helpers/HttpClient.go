package helpers

import (
	"net/http"
	"net/url"
	"os"
)

func RequestClient(input *http.Client) *http.Client {
	client := &http.Client{}

	if os.Getenv("ENV") != "production" {
		proxyUrl, _ := url.Parse("http://10.126.111.123:4480")
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	}

	return client
}
