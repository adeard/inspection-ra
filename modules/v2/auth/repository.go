package auth_v2

import (
	"encoding/json"
	"inspection-ra/domain"
	"inspection-ra/helpers"
	"os"

	"gorm.io/gorm"
)

type Repository interface {
	SignIn(auth domain.Auth) (domain.NewUserResponse, error)
	GetLogged(token string) (domain.NewDetailUserResponse, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SignIn(auth domain.Auth) (domain.NewUserResponse, error) {
	input := struct {
		Username string `json:"user_name"`
		Password string `json:"password"`
	}{
		auth.Username,
		auth.Password,
	}

	url := os.Getenv("USER_SERVICE_URL") + "/api/user/login"

	body, _ := helpers.SendRequest("POST", url, input)

	var returnLogin domain.NewUserResponse
	if err := json.Unmarshal(body, &returnLogin); err != nil {
		return returnLogin, err
	}

	return returnLogin, nil
}

func (r *repository) GetLogged(token string) (domain.NewDetailUserResponse, error) {

	result, _ := helpers.GetDetailUser(token)

	return result, nil
}
