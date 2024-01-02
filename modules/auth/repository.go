package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"inspection-ra/domain"
	"inspection-ra/helpers"
	"io"
	"net/http"
	"net/url"
	"os"

	"gorm.io/gorm"
)

type Repository interface {
	AuthTest() (domain.Auth, error)
	SignIn(auth domain.Auth) (domain.AuthData, error)
	GetLogged(token string) (domain.AuthLoggedData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AuthTest() (domain.Auth, error) {
	auth := domain.Auth{
		Username: "test username",
		Password: "test password",
	}

	return auth, nil
}

func (r *repository) SignIn(auth domain.Auth) (domain.AuthData, error) {
	authInput, err := json.Marshal(auth)
	if err != nil {
		return domain.AuthData{}, err
	}

	urladdr := os.Getenv("USER_SERVICE_URL") + "/auth/sign_in"
	fmt.Println("URL:>", urladdr)

	var jsonStr = []byte(authInput)
	req, err := http.NewRequest("POST", urladdr, bytes.NewBuffer(jsonStr))
	if err != nil {
		return domain.AuthData{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	// proxyUrl, _ := url.Parse("http://10.126.111.123:4480")
	// client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		helpers.LogInit(err.Error())
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result domain.AuthData
	if err := json.Unmarshal(body, &result); err != nil {
		return result, err
	}

	return result, nil
}

func (r *repository) GetLogged(token string) (domain.AuthLoggedData, error) {

	urladdr := os.Getenv("USER_SERVICE_URL") + "/api/user/user_profile/front/self"
	fmt.Println("URL:>", urladdr)

	req, err := http.NewRequest("GET", urladdr, nil)
	if err != nil {
		return domain.AuthLoggedData{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	proxyUrl, _ := url.Parse("http://10.126.111.123:4480")
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var parsed domain.AuthLoggedParsedData

	if err := json.Unmarshal(body, &parsed); err != nil {
		return domain.AuthLoggedData{}, err
	}

	result := domain.AuthLoggedData{
		Id:       parsed.Id,
		Code:     parsed.Code,
		Username: parsed.Username,
		Email:    parsed.Email,
		Fullname: parsed.Employees[0].Fullname,
		Company:  parsed.Company,
	}

	return result, nil
}
