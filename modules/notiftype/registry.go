package notiftype

import "gorm.io/gorm"

func NotifTypeRegistry(db *gorm.DB) Service {
	notifTypeRepository := NewRepository(db)
	notifTypeService := NewService(notifTypeRepository)

	return notifTypeService
}
