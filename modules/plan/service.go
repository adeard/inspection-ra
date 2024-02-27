package plan

import (
	"inspection-ra/domain"
)

type Service interface {
	Store(input domain.PlanRequest) (domain.PlanRequest, error)
	Insert(input []domain.PlanRequest) (string, error)
	GetAll(input domain.PlanRequest) ([]domain.PlanData, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(input domain.PlanRequest) ([]domain.PlanData, error) {

	plan, err := s.repository.FindAll(input)

	return plan, err
}

func (s *service) Store(input domain.PlanRequest) (domain.PlanRequest, error) {

	planData, err := s.repository.Store(input)

	return planData, err
}

func (s *service) Insert(input []domain.PlanRequest) (string, error) {
	var err error
	var result string
	existPlanIds := []int32{}
	insertedData := []domain.PlanRequest{}

	for index, plan := range input {
		filterCheck := domain.PlanRequest{
			CompanyCode:    plan.CompanyCode,
			Ba:             plan.Ba,
			RunningAccount: plan.RunningAccount,
			PlanDate:       plan.PlanDate,
		}

		check, err := s.repository.GetDetail(filterCheck)
		if err == nil {
			if input[index].InspectDate == "" {
				continue
			}
			existPlanIds = append(existPlanIds, check.Id)
		}

		insertedData = append(insertedData, plan)

		if len(insertedData) > 500 {
			if len(existPlanIds) > 0 {
				s.repository.DeleteBatch(existPlanIds)
			}

			result, err = s.repository.Insert(insertedData)
			if err != nil {
				return result, err
			} else {
				existPlanIds = []int32{}
				insertedData = []domain.PlanRequest{}
			}
		}
	}

	if len(existPlanIds) > 0 {
		s.repository.DeleteBatch(existPlanIds)
	}

	if len(insertedData) > 0 {
		result, err = s.repository.Insert(insertedData)
	}

	return result, err
}
