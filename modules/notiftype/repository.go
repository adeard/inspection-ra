package notiftype

import (
	"inspection-ra/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(input domain.NotifTypeRequest) ([]domain.NotifTypeData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(input domain.NotifTypeRequest) ([]domain.NotifTypeData, error) {
	var notifType []domain.NotifTypeData

	q := r.db

	if input.Notif_Type != "" {
		q = q.Where("Notif_Type =  ? ", input.Notif_Type)
	}

	if input.Plant != "" {
		q = q.Where("Plant =  ? ", input.Plant)
	}

	if input.Notif_Desc != "" {
		q = q.Where("Notif_Desc =  ? ", input.Notif_Desc)
	}

	if input.Application != "" {
		q = q.Where("Application =  ? ", input.Application)
	}

	if input.Catalog_Profile != "" {
		q = q.Where("Catalog_Profile =  ? ", input.Catalog_Profile)
	}

	err := q.Find(&notifType).Error

	return notifType, err
}
