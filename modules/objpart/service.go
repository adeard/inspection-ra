package objpart

import (
	"inspection-ra/domain"
)

type Service interface {
	GetAll() ([]domain.ObjPartDataResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.ObjPartDataResponse, error) {
	objPartDataResp := []domain.ObjPartDataResponse{}

	objPart, err := s.repository.FindAll()
	if err != nil {
		return objPartDataResp, err
	}

	getAllDamage, err := s.repository.FindAllDamage()
	if err != nil {
		return objPartDataResp, err
	}

	getAllCause, err := s.repository.FindAllCause()
	if err != nil {
		return objPartDataResp, err
	}

	for _, v := range objPart {
		damageList := []domain.DamageDataResponse{}

		for _, damageDataValue := range getAllDamage {
			damageData := domain.DamageDataResponse{}
			causeList := []domain.CauseDataResponse{}
			if damageDataValue.ObjPartCode == v.ObjPartCode && damageDataValue.ZinspecVehicleTypeId == v.ZinspecVehicleTypeId {
				damageData.DamageData = damageDataValue

				for _, causeDataValue := range getAllCause {
					causeData := domain.CauseDataResponse{}
					if causeDataValue.ObjPartCode == damageDataValue.ObjPartCode && causeDataValue.DamageCode == damageDataValue.DamageCode && causeDataValue.ZinspecVehicleTypeId == damageData.ZinspecVehicleTypeId {
						causeData.CauseData = causeDataValue
						causeList = append(causeList, causeData)
					}
				}

				damageData.Causes = causeList
				damageList = append(damageList, damageData)
			}
		}

		responseTemp := domain.ObjPartDataResponse{
			Id:              v.Id,
			ObjPartRequest:  v.ObjPartRequest,
			VehicleTypeCode: v.VehicleTypeData.Code,
			Damages:         damageList,
		}

		objPartDataResp = append(objPartDataResp, responseTemp)
	}

	return objPartDataResp, err
}
