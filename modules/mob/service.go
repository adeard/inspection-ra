package mob

import "inspection-ra/domain"

type Service interface {
	GetAll() ([]domain.MobData, error)
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
