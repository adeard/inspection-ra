package mob

import "inspection-ra/domain"

type Service interface {
	GetAll() ([]domain.MobData, error)
	Insert(input []domain.MobRequest) ([]domain.MobRequest, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.MobData, error) {

	mob, err := s.repository.FindAll()

	return mob, err
}

func (s *service) Insert(input []domain.MobRequest) ([]domain.MobRequest, error) {

	mobData, err := s.repository.Insert(input)

	return mobData, err
}
