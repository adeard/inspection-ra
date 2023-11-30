package plan

import "inspection-ra/domain"

type Service interface {
	GetAll() ([]domain.PlanData, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.PlanData, error) {

	plan, err := s.repository.FindAll()

	return plan, err
}
