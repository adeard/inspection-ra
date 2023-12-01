package plan

import (
	"inspection-ra/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.PlanData, error)
	Store(input domain.PlanRequest) (domain.PlanRequest, error)
	Insert(input []domain.PlanRequest) ([]domain.PlanRequest, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.PlanData, error) {
	var plan []domain.PlanData
	err := r.db.Find(&plan).Error

	return plan, err
}

func (r *repository) Store(planData domain.PlanRequest) (domain.PlanRequest, error) {
	err := r.db.Create(&planData).Error

	return planData, err
}

func (r *repository) Insert(planData []domain.PlanRequest) ([]domain.PlanRequest, error) {
	err := r.db.Create(&planData).Error

	return planData, err
}
