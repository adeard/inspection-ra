package mob

import (
	"inspection-ra/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.MobData, error)
	Insert(input []domain.MobRequest) ([]domain.MobRequest, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.MobData, error) {
	var mob []domain.MobData
	err := r.db.Find(&mob).Error

	return mob, err
}

func (r *repository) Insert(mobData []domain.MobRequest) ([]domain.MobRequest, error) {
	err := r.db.Create(&mobData).Error

	return mobData, err
}
