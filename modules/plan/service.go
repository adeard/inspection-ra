package plan

import "inspection-ra/domain"

type Service interface {
	GetAll() ([]domain.PlanData, error)
	Store(input domain.PlanRequest) (domain.PlanRequest, error)
	Insert(input []domain.PlanRequest) ([]domain.PlanRequest, error)
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

func (s *service) Store(input domain.PlanRequest) (domain.PlanRequest, error) {

	planData, err := s.repository.Store(input)

	return planData, err
}

func (s *service) Insert(input []domain.PlanRequest) ([]domain.PlanRequest, error) {

	planData, err := s.repository.Insert(input)

	return planData, err
}
