package runacct

import "inspection-ra/domain"

type Service interface {
	GetAll() ([]domain.RunAcctData, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.RunAcctData, error) {

	runAcct, err := s.repository.FindAll()

	return runAcct, err
}
