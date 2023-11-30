package plan

import "gorm.io/gorm"

func PlanRegistry(db *gorm.DB) Service {
	planRepository := NewRepository(db)
	planService := NewService(planRepository)

	return planService
}
