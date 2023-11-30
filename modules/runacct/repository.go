package runacct

import (
	"inspection-ra/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.RunAcctData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.RunAcctData, error) {
	var runAcct []domain.RunAcctData
	err := r.db.Find(&runAcct).Error

	return runAcct, err
}
