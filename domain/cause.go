package domain

type CauseResponse struct {
	Data        any    `json:"data"`
	Message     string `json:"message"`
	ElapsedTime string `json:"elapsed_time"`
}

type CauseRequest struct {
	Catalog              string `json:"catalog" gorm:"column:catalog;"`
	VehicleType          string `json:"vehicle_type" gorm:"column:vehicle_type;"`
	ObjPartCode          string `json:"objpart_code" gorm:"column:objpart_code;"`
	DamageCode           string `json:"damage_code" gorm:"column:damage_code;"`
	CauseCode            string `json:"cause_code" gorm:"column:causes_code;"`
	Cause                string `json:"cause" gorm:"column:causes;"`
	ZinspecVehicleTypeId int32  `json:"zinspec_vehicletype_id" gorm:"column:zinspec_vehicletype_id;"`
}

type CauseData struct {
	CauseRequest
}

type CauseDataResponse struct {
	CauseData
}

func (CauseData) TableName() string {
	return "zinspec_causes"
}
