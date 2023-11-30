package objpart

import (
	"inspection-ra/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.ObjPartData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.ObjPartData, error) {
	var objPart []domain.ObjPartData
	err := r.db.Find(&objPart).Error

	return objPart, err
}
