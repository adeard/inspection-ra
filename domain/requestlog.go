package domain

type RequestLogResponse struct {
	Data        any    `json:"data"`
	Message     string `json:"message"`
	ElapsedTime string `json:"elapsed_time"`
}

type RequestLogRequest struct {
	StatusCode int    `json:"status_code" gorm:"column:status_code;"`
	Source     string `json:"source" gorm:"column:source;"`
	Request    string `json:"request" gorm:"column:request;"`
	Response   string `json:"response" gorm:"column:response;"`
	CreatedAt  string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt  string `json:"updated_at" gorm:"column:updated_at;"`
	Method     string `json:"method" gorm:"column:method;"`
	ClientIp   string `json:"client_ip" gorm:"column:client_ip;"`
	UserAgent  string `json:"user_agent" gorm:"column:user_agent;"`
	Duration   string `json:"duration" gorm:"column:duration;"`
}

type RequestLogData struct {
	Id int32 `json:"id" gorm:"column:id;"`
	RequestLogRequest
}

func (RequestLogRequest) TableName() string {
	return "RequestLog"
}
