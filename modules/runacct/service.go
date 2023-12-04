package runacct

import "inspection-ra/domain"

type Service interface {
	GetAll(input domain.RunAcctRequest) ([]domain.RunAcctData, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(input domain.RunAcctRequest) ([]domain.RunAcctData, error) {

	runAcct, err := s.repository.FindAll(input)

	return runAcct, err
}
