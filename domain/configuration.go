package domain

type ConfigurationResponse struct {
	Data        any    `json:"data"`
	Message     string `json:"message"`
	ElapsedTime string `json:"elapsed_time"`
}

type ConfigurationRequest struct {
	ValidFrom    string `json:"valid_from" gorm:"column:valid_from;"`
	Interval     int32  `json:"interval" gorm:"column:interval;"`
	IntervalType string `json:"interval_type" gorm:"column:interval_type;"`
	LastWeek     int32  `json:"last_week" gorm:"column:last_week;"`
	IsActive     bool   `json:"is_active" gorm:"column:is_active;"`
	PlanDate     string `json:"plan_date" gorm:"column:plan_date;"`
}

type ConfigurationData struct {
	Id int32 `json:"id" gorm:"column:id;"`
	ConfigurationRequest
}

func (ConfigurationRequest) TableName() string {
	return "zinspec_configuration"
}
