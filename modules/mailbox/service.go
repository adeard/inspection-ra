package mailbox

import (
	"inspection-ra/domain"
	"time"
)

type Service interface {
	Store(input domain.MailboxRequest) (domain.MailboxRequest, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Store(input domain.MailboxRequest) (domain.MailboxRequest, error) {

	location, _ := time.LoadLocation("Local")

	currentDateTime := time.Now().In(location).Format("2006-01-02 15:04:05")

	input.CreatedAt = currentDateTime
	input.UpdatedAt = currentDateTime

	log, err := s.repository.Store(input)
	if err != nil {
		return domain.MailboxRequest{}, err
	}

	return log, err
}
