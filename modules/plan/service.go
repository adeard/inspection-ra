package plan

import (
	"inspection-ra/domain"
)

type Service interface {
	Store(input domain.PlanRequest) (domain.PlanRequest, error)
	Insert(input []domain.PlanRequest) (string, error)
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

func (s *service) Store(input domain.PlanRequest) (domain.PlanRequest, error) {

	planData, err := s.repository.Store(input)

	return planData, err
}

func (s *service) Insert(input []domain.PlanRequest) (string, error) {

	existPlanIds := []int32{}

	for _, plan := range input {
		filterCheck := domain.PlanRequest{
			CompanyCode:    plan.CompanyCode,
			Ba:             plan.Ba,
			RunningAccount: plan.RunningAccount,
			PlanDate:       plan.PlanDate,
		}

		check, err := s.repository.GetDetail(filterCheck)
		if err == nil {
			existPlanIds = append(existPlanIds, check.Id)
		}
	}

	if len(existPlanIds) > 0 {
		s.repository.DeleteBatch(existPlanIds)
	}

	result, err := s.repository.Insert(input)

	return result, err
}
