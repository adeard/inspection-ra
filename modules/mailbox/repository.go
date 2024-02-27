package mailbox

import (
	"inspection-ra/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.MailboxRequest) (domain.MailboxRequest, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Store(input domain.MailboxRequest) (domain.MailboxRequest, error) {
	err := r.db.Create(&input).Error

	return input, err
}
