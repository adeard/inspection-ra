package configuration

import (
	"inspection-ra/domain"

	"gorm.io/gorm"
)

type Repository interface {
	UpdateById(id int32, updateData map[string]interface{}) error
	GetDetail(input domain.ConfigurationRequest) (domain.ConfigurationData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetDetail(input domain.ConfigurationRequest) (domain.ConfigurationData, error) {
	var config domain.ConfigurationData

	q := r.db.Debug()

	if input.ValidFrom != "" {
		q = q.Where("valid_from =  ? ", input.ValidFrom)
	}

	if input.Interval > 0 {
		q = q.Where("interval =  ? ", input.Interval)
	}

	if input.IntervalType != "" {
		q = q.Where("interval_type =  ? ", input.IntervalType)
	}

	if input.LastWeek > 0 {
		q = q.Where("last_week =  ? ", input.LastWeek)
	}

	if input.IsActive {
		q = q.Where("is_active =  ? ", input.IsActive)
	}

	if input.PlanDate != "" {
		q = q.Where("plan_date =  ? ", input.PlanDate)
	}

	err := q.First(&config).Error

	return config, err
}

func (r *repository) UpdateById(id int32, updateData map[string]interface{}) error {
	err := r.db.Debug().
		Table("zinspec_configuration").
		Where("id = ?", id).
		Updates(updateData).Error

	return err
}
