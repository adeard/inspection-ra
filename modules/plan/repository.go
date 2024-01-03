package plan

import (
	"inspection-ra/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.PlanRequest) (domain.PlanRequest, error)
	Insert(input []domain.PlanRequest) (string, error)
	FindAll(input domain.PlanRequest) ([]domain.PlanData, error)
	GetById(id int32) (domain.PlanData, error)
	GetDetail(input domain.PlanRequest) (domain.PlanData, error)
	DeleteBatch(ids []int32) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(input domain.PlanRequest) ([]domain.PlanData, error) {
	var plan []domain.PlanData

	q := r.db.Debug().Table("zinspec_plan")

	if input.Ba != "" {
		q = q.Where("ba = ?", input.Ba)
	}

	if input.Week > 0 {
		q = q.Where("week = ?", input.Week)
	}

	if input.PlanDate != "" {
		q = q.Where("plan_date = ?", input.PlanDate)
	}

	if input.CompanyCode != "" {
		q = q.Where("company_code = ?", input.CompanyCode)
	}

	err := q.Find(&plan).Error

	return plan, err
}

func (r *repository) Store(planData domain.PlanRequest) (domain.PlanRequest, error) {
	err := r.db.Create(&planData).Error

	return planData, err
}

func (r *repository) Insert(planData []domain.PlanRequest) (string, error) {
	err := r.db.Create(&planData).Error

	return "success", err
}

func (r *repository) DeleteBatch(ints []int32) (string, error) {
	plan := domain.PlanData{}

	err := r.db.Delete(&plan, ints).Error

	return "success", err
}

func (r *repository) GetById(id int32) (domain.PlanData, error) {
	plan := domain.PlanData{}

	err := r.db.First(&plan, id).Error

	return plan, err
}

func (r *repository) GetDetail(input domain.PlanRequest) (domain.PlanData, error) {
	plan := domain.PlanData{}

	q := r.db.Debug().Table("zinspec_plan")

	if input.Ba != "" {
		q = q.Where("ba = ?", input.Ba)
	}

	if input.CompanyCode != "" {
		q = q.Where("company_code = ?", input.CompanyCode)
	}

	if input.RunningAccount != "" {
		q = q.Where("running_account = ?", input.RunningAccount)
	}

	if input.Status != "" {
		q = q.Where("status = ?", input.Status)
	}

	if input.PlanDate != "" {
		q = q.Where("plan_date = ?", input.PlanDate)
	}

	if input.InspectDate != "" {
		q = q.Where("inspect_date = ?", input.InspectDate)
	}

	if input.InspectTime != "" {
		q = q.Where("inspect_time = ?", input.InspectTime)
	}

	if input.Week != 0 {
		q = q.Where("week = ?", input.Week)
	}

	err := q.First(&plan).Error

	return plan, err
}
