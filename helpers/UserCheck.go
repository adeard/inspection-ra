package helpers

import (
	"encoding/json"
	"inspection-ra/domain"
	"io/ioutil"
	"net/http"
	"os"
)

func CheckUserAccess(token string) (string, error) {
	urladdr := "https://user-api-simp.azurewebsites.net/api/APP782231326/verify_save"
	req, _ := http.NewRequest("GET", urladdr, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Accept", "application/json")

	// Send req using http Client
	client := RequestClient(&http.Client{})
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

func GetDetailUser(token string) (domain.NewDetailUserResponse, error) {
	urladdr := os.Getenv("NEW_USER_SERVICE_URL") + "/api/user/detail_by_name_application?name_application=InspectionRA"
	req, _ := http.NewRequest("GET", urladdr, nil)
	req.Header.Add("authenticationToken", token)
	req.Header.Add("Accept", "application/json")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return domain.NewDetailUserResponse{}, err
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	result := domain.NewDetailUserResponse{}
	json.Unmarshal(data, &result)

	return result, err
}
