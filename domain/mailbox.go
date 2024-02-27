package domain

type MailboxResponse struct {
	Data        any    `json:"data"`
	Message     string `json:"message"`
	ElapsedTime string `json:"elapsed_time"`
}

type MailboxRequest struct {
	StatusCode int    `json:"status_code" gorm:"column:status_code;"`
	Source     string `json:"source" gorm:"column:source;"`
	Request    string `json:"request" gorm:"column:request;"`
	Response   string `json:"response" gorm:"column:response;"`
	CreatedAt  string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt  string `json:"updated_at" gorm:"column:updated_at;"`
}

type MailboxData struct {
	Id int32 `json:"id" gorm:"column:id;"`
	MailboxRequest
}

func (MailboxRequest) TableName() string {
	return "zinspec_mailbox"
}
