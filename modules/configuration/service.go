package configuration

import (
	"inspection-ra/domain"
	"time"

	"github.com/jinzhu/now"
)

type Service interface {
	CheckConfig(input domain.ConfigurationRequest) (domain.ConfigurationData, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CheckConfig(input domain.ConfigurationRequest) (domain.ConfigurationData, error) {

	maxWeekYear := int32(52)

	_, week := time.Now().ISOWeek()

	input.IsActive = true

	configData, err := s.repository.GetDetail(input)

	maxWeek := configData.LastWeek + (configData.Interval - 1)

	if week > int(maxWeek) || maxWeek > maxWeekYear {
		updateData := map[string]interface{}{
			"last_week": int32(week),
			"plan_date": now.BeginningOfWeek().AddDate(0, 0, 1).Format("2006-01-02"),
		}

		err = s.repository.UpdateById(configData.Id, updateData)
		if err == nil {
			configData.LastWeek = int32(week)
			configData.PlanDate = now.BeginningOfWeek().AddDate(0, 0, 1).Format("2006-01-02")
		}
	}

	return configData, err
}
