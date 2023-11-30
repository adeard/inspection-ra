package plan

import (
	"inspection-ra/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.PlanData, error)
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
