package auth_v2

import (
	"errors"
	"fmt"
	"inspection-ra/domain"
	"inspection-ra/helpers"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	SignIn(authinput domain.AuthRequest) (domain.AuthData, error)
	GetLogged(token string) (domain.NewDetailUserResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SignIn(authinput domain.AuthRequest) (domain.AuthData, error) {

	auth, err := s.repository.SignIn(domain.Auth(authinput))
	if err != nil {
		return domain.AuthData{}, err
	}

	loggedData, err := s.repository.GetLogged(auth.Datas)
	if err != nil {
		return domain.AuthData{}, err
	}

	objek, ok := loggedData.Objek.(map[string]interface{})
	if !ok {
		fmt.Println("d is not interface")
	}

	applications := objek["Applications"].([]interface{})
	if applications[0].(float64) == 0 {
		err = errors.New("user not registered")
		return domain.AuthData{}, err
	}

	userId := objek["Id"].(float64)
	fullname := objek["FullName"].(string)

	generateJWT, err := helpers.GenerateJWT(int(userId), fullname, auth.Datas)
	if err != nil {
		return domain.AuthData{}, err
	}

	result := domain.AuthData{
		Token:   generateJWT,
		Message: auth.Message,
	}

	return result, err
}

func (s *service) GetLogged(token string) (domain.NewDetailUserResponse, error) {

	parsedToken, err := helpers.ParseToken(token)
	if err != nil {
		return domain.NewDetailUserResponse{}, err
	}

	claims, _ := parsedToken.Claims.(jwt.MapClaims)

	authToken := claims["auth_token"].(string)

	loggedData, err := s.repository.GetLogged(authToken)

	return loggedData, err
}
