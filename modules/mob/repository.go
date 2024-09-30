package mob

import (
	"inspection-ra/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Insert(input []domain.MobRequest) (string, error)
	FindAll(input domain.MobRequest) ([]domain.MobData, error)
	GetDetail(input domain.MobRequest) (domain.MobData, error)
	DeleteBatch(ids []int32) (string, error)
	StoreMobItemDamaged(input domain.MobItemDamagedRequest) (domain.MobItemDamagedRequest, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(input domain.MobRequest) ([]domain.MobData, error) {
	var mob []domain.MobData

	q := r.db.Debug().Table("zinspec_mob")

	if input.Ba != "" {
		q = q.Where("ba = ?", input.Ba)
	}

	if input.PlanDate != "" {
		q = q.Where("plan_date = ?", input.PlanDate)
	}

	err := q.Find(&mob).Error

	return mob, err
}

func (r *repository) GetDetail(input domain.MobRequest) (domain.MobData, error) {
	var mob domain.MobData

	q := r.db.Debug().Table("zinspec_mob")

	if input.Ba != "" {
		q = q.Where("ba = ?", input.Ba)
	}

	if input.PlanDate != "" {
		q = q.Where("plan_date = ?", input.PlanDate)
	}

	if input.NoInspec != "" {
		q = q.Where("no_inspec = ?", input.NoInspec)
	}

	err := q.First(&mob).Error

	return mob, err
}

func (r *repository) Insert(mobData []domain.MobRequest) (string, error) {
	tx := r.db.Begin()

	err := tx.Create(&mobData).Error
	if err != nil {
		tx.Rollback()

		return "error", err
	}

	tx.Commit()

	return "success", err
}

func (r *repository) DeleteBatch(ints []int32) (string, error) {
	mob := domain.MobData{}

	err := r.db.Delete(&mob, ints).Error

	return "success", err
}

func (r *repository) StoreMobItemDamaged(input domain.MobItemDamagedRequest) (domain.MobItemDamagedRequest, error) {
	err := r.db.Create(&input).Error

	return input, err
}
