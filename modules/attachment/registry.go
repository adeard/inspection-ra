package attachment

import "gorm.io/gorm"

func AttachmentRegistry(db *gorm.DB) Service {
	attachmentRepository := NewRepository(db)
	attachmentService := NewService(attachmentRepository)

	return attachmentService
}
