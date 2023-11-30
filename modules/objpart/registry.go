package objpart

import "gorm.io/gorm"

func ObjPartRegistry(db *gorm.DB) Service {
	objPartRepository := NewRepository(db)
	objPartService := NewService(objPartRepository)

	return objPartService
}
