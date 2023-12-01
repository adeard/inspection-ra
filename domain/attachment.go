package domain

import "mime/multipart"

type AttachmentData struct {
	Id int32 `json:"id" form:"id" gorm:"column:id"`
	AttachmentRequest
}

type AttachmentRequest struct {
	NoInspec      string                `json:"no_inspec" form:"no_inspec" gorm:"column:no_inspec"`
	ZinspecMobId  int32                 `json:"zinspec_mob_id" form:"zinspec_mob_id" gorm:"column:zinspec_mob_id"`
	ImageCategory string                `json:"image_category" form:"image_category" gorm:"column:image_category"`
	Filename      string                `json:"filename" form:"filename" gorm:"column:filename"`
	File          *multipart.FileHeader `json:"file" form:"file" binding:"required"`
	CreatedAt     string                `json:"created_at" form:"created_at" gorm:"column:created_at"`
}

func (AttachmentRequest) TableName() string {
	return "zinspec_attachment"
}

type AttachmentResponse struct {
	Data        any    `json:"data"`
	Message     string `json:"message"`
	ElapsedTime string `json:"elapsed_time"`
}
