package mob

import "gorm.io/gorm"

func MobRegistry(db *gorm.DB) Service {
	mobRepository := NewRepository(db)
	mobService := NewService(mobRepository)

	return mobService
}
