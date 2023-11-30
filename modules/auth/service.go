package auth

import "inspection-ra/domain"

type Service interface {
	AuthTest() (domain.Auth, error)
	SignIn(authinput domain.AuthRequest) (domain.AuthData, error)
	GetLogged(token string) (domain.AuthLoggedData, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) AuthTest() (domain.Auth, error) {

	auth, err := s.repository.AuthTest()

	return auth, err
}

func (s *service) SignIn(authinput domain.AuthRequest) (domain.AuthData, error) {

	authinput.Ldap = true

	auth, err := s.repository.SignIn(domain.Auth(authinput))

	return auth, err
}

func (s *service) GetLogged(token string) (domain.AuthLoggedData, error) {

	loggedData, err := s.repository.GetLogged(token)

	return loggedData, err
}
