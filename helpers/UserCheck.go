package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CheckUserAccess(token string) (string, error) {
	urladdr := "https://user-api-simp.azurewebsites.net/api/APP782231326/verify_save"
	req, _ := http.NewRequest("GET", urladdr, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Accept", "application/json")

	// Send req using http Client
	// proxyUrl, _ := url.Parse("http://10.126.111.123:4480")
	// client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err.Error(), err
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	var message struct {
		Message string
	}

	json.Unmarshal(data, &message)

	return message.Message, err
}
