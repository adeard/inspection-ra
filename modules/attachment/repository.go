package attachment

import (
	"inspection-ra/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.AttachmentData) (domain.AttachmentData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Store(attachment domain.AttachmentData) (domain.AttachmentData, error) {
	err := r.db.Create(&attachment).Error

	return attachment, err
}
