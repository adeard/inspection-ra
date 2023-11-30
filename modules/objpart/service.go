package objpart

import "inspection-ra/domain"

type Service interface {
	GetAll() ([]domain.ObjPartData, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.ObjPartData, error) {

	objPart, err := s.repository.FindAll()

	return objPart, err
}
