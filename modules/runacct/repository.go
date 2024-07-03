package runacct

import (
	"inspection-ra/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.RunAcctRequest) ([]domain.RunAcctData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(input domain.RunAcctRequest) ([]domain.RunAcctData, error) {
	var runAcct []domain.RunAcctData

	q := r.db.Debug()

	if input.Estate != "" {
		q = q.Where("estate =  ? ", input.Estate)
	}

	if input.Anln1 != "" {
		q = q.Where("anln1 =  ? ", input.Anln1)
	}

	if input.StatusFlag != "" {
		q = q.Where("status_flag =  ? ", input.StatusFlag)
	}

	err := q.Find(&runAcct).Error

	return runAcct, err
}
