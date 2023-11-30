package vehicletype

import "gorm.io/gorm"

func VehicleTypeRegistry(db *gorm.DB) Service {
	vehicleTypeRepository := NewRepository(db)
	vehicleTypeService := NewService(vehicleTypeRepository)

	return vehicleTypeService
}
