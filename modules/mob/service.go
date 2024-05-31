package mob

import (
	"inspection-ra/domain"
	"inspection-ra/helpers"
)

type Service interface {
	GetAll(input domain.MobRequest) ([]domain.MobData, error)
	Insert(input []domain.MobRequest) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(input domain.MobRequest) ([]domain.MobData, error) {

	mob, err := s.repository.FindAll(input)

	return mob, err
}

func (s *service) Insert(input []domain.MobRequest) (string, error) {
	var err error
	var result string
	existMobIds := []int32{}
	insertedData := []domain.MobRequest{}

	for _, mobData := range input {
		filterCheck := domain.MobRequest{
			NoInspec: mobData.NoInspec,
		}

		check, err := s.repository.GetDetail(filterCheck)
		if err == nil {
			existMobIds = append(existMobIds, check.Id)
		}

		mobData.SyncDate = helpers.GetCurrentDateTime()
		insertedData = append(insertedData, mobData)

		if len(insertedData) > 70 {
			if len(existMobIds) > 0 {
				s.repository.DeleteBatch(existMobIds)
			}

			result, err = s.repository.Insert(insertedData)
			if err != nil {
				return result, err
			} else {
				existMobIds = []int32{}
				insertedData = []domain.MobRequest{}
			}
		}
	}

	if len(existMobIds) > 0 {
		s.repository.DeleteBatch(existMobIds)
	}

	if len(insertedData) > 0 {
		result, err = s.repository.Insert(insertedData)
	}

	return result, err
}
