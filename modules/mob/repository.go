package mob

import (
	"inspection-ra/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.MobRequest) ([]domain.MobData, error)
	Insert(input []domain.MobRequest) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(input domain.MobRequest) ([]domain.MobData, error) {
	var mob []domain.MobData

	q := r.db.Debug().Table("zinspec_mob")

	if input.Ba != "" {
		q = q.Where("ba = ?", input.Ba)
	}

	if input.PlanDate != "" {
		q = q.Where("plan_date = ?", input.PlanDate)
	}

	err := q.Find(&mob).Error

	return mob, err
}

func (r *repository) Insert(mobData []domain.MobRequest) (string, error) {
	err := r.db.Create(&mobData).Error

	return "success", err
}
