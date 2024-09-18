package objpart

import (
	"inspection-ra/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.ObjPartData, error)
	FindAllCause() ([]domain.CauseData, error)
	FindAllDamage() ([]domain.DamageData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.ObjPartData, error) {
	var objPart []domain.ObjPartData
	err := r.db.Preload("VehicleTypeData").Find(&objPart).Error

	return objPart, err
}

func (r *repository) FindAllDamage() ([]domain.DamageData, error) {
	var damage []domain.DamageData
	err := r.db.Find(&damage).Error

	return damage, err
}

func (r *repository) FindAllCause() ([]domain.CauseData, error) {
	var cause []domain.CauseData
	err := r.db.Find(&cause).Error

	return cause, err
}
