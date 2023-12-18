package configuration

import "gorm.io/gorm"

func ConfigurationRegistry(db *gorm.DB) Service {
	configurationRepository := NewRepository(db)
	configurationService := NewService(configurationRepository)

	return configurationService
}
