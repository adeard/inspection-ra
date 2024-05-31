package auth_v2

import (
	"inspection-ra/domain"
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

	result := domain.AuthData{
		Token:   auth.Datas,
		Message: auth.Message,
	}

	return result, err
}

func (s *service) GetLogged(token string) (domain.NewDetailUserResponse, error) {

	loggedData, err := s.repository.GetLogged(token)

	return loggedData, err
}
