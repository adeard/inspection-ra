package notiftype

import "inspection-ra/domain"

type Service interface {
	GetAll(input domain.NotifTypeRequest) ([]domain.NotifTypeData, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(input domain.NotifTypeRequest) ([]domain.NotifTypeData, error) {

	runAcct, err := s.repository.FindAll(input)

	return runAcct, err
}
