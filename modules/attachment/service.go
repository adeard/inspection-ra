package attachment

import (
	"fmt"
	"inspection-ra/domain"
	"os"
	"time"
)

type Service interface {
	Store(input domain.AttachmentRequest, destination string) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Store(input domain.AttachmentRequest, destination string) (string, error) {
	attachmentData := domain.AttachmentData{
		NoInspec:      input.NoInspec,
		ZinspecMobId:  input.ZinspecMobId,
		ImageCategory: input.ImageCategory,
		Filename:      input.Filename,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	checkAttachmentData, err := s.repository.GetDetail(input)
	if err == nil {
		s.repository.DeleteById(checkAttachmentData.Id)
		err := os.Remove(destination + checkAttachmentData.Filename)
		if err != nil {
			fmt.Println(err)
		}
	}

	attachment, err := s.repository.Store(attachmentData)

	return attachment, err
}
