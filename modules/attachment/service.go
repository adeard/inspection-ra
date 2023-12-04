package attachment

import (
	"inspection-ra/domain"
	"time"
)

type Service interface {
	Store(input domain.AttachmentRequest) (domain.AttachmentData, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Store(input domain.AttachmentRequest) (domain.AttachmentData, error) {
	attachmentData := domain.AttachmentData{
		NoInspec:      input.NoInspec,
		ZinspecMobId:  input.ZinspecMobId,
		ImageCategory: input.ImageCategory,
		Filename:      input.Filename,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}
	attachment, err := s.repository.Store(attachmentData)

	return attachment, err
}
