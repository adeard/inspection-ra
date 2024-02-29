package attachment

import (
	"inspection-ra/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.AttachmentData) (string, error)
	GetDetail(input domain.AttachmentRequest) (domain.AttachmentData, error)
	DeleteById(id int32) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Store(attachment domain.AttachmentData) (string, error) {
	err := r.db.Create(&attachment).Error

	return "success", err
}

func (r *repository) GetDetail(input domain.AttachmentRequest) (domain.AttachmentData, error) {
	var attachmentData domain.AttachmentData

	q := r.db.Table("zinspec_attachment")

	if input.NoInspec != "" {
		q = q.Where("no_inspec =  ? ", input.NoInspec)
	}

	if input.ImageCategory != "" {
		q = q.Where("image_category =  ? ", input.ImageCategory)
	}

	err := q.First(&attachmentData).Error

	return attachmentData, err
}

func (r *repository) DeleteById(id int32) (string, error) {

	err := r.db.Delete(&domain.AttachmentData{}, id).Error

	return "success", err
}
