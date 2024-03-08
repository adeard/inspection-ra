package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"inspection-ra/domain"
	"inspection-ra/helpers"
	"io"
	"net/http"
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

	client := helpers.RequestClient(&http.Client{})
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

	client := helpers.RequestClient(&http.Client{})
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
		Company:  parsed.Company,
	}

	if parsed.Fullname != "" {
		result.Fullname = parsed.Fullname
	}

	if len(parsed.Employees) > 0 {
		result.Fullname = parsed.Employees[0].Fullname
	}

	return result, nil
}
