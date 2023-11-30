package runacct

import "gorm.io/gorm"

func RunAcctRegistry(db *gorm.DB) Service {
	runAcctRepository := NewRepository(db)
	runAcctService := NewService(runAcctRepository)

	return runAcctService
}
