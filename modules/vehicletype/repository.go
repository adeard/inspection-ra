package vehicletype

import (
	"inspection-ra/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.VehicleTypeData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.VehicleTypeData, error) {
	var vehicleType []domain.VehicleTypeData
	err := r.db.Find(&vehicleType).Error

	return vehicleType, err
}
