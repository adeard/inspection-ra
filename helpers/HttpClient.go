package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func SendRequest(method string, url string, input any) ([]byte, error) {
	authInput, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	fmt.Println("URL:>", url)

	var req *http.Request

	if method == "POST" {
		jsonStr := []byte(authInput)
		req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	} else {
		req, err = http.NewRequest("GET", url, nil)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		LogInit(err.Error())
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return body, nil
}
