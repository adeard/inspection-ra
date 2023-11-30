package vehicletype

import "inspection-ra/domain"

type Service interface {
	GetAll() ([]domain.VehicleTypeData, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.VehicleTypeData, error) {

	vehicletype, err := s.repository.FindAll()

	return vehicletype, err
}
