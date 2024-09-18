package domain

type DamageResponse struct {
	Data        any    `json:"data"`
	Message     string `json:"message"`
	ElapsedTime string `json:"elapsed_time"`
}

type DamageRequest struct {
	Catalog              string `json:"catalog" gorm:"column:catalog;"`
	VehicleType          string `json:"vehicle_type" gorm:"column:vehicle_type;"`
	ObjPartCode          string `json:"objpart_code" gorm:"column:objpart_code;"`
	DamageCode           string `json:"damage_code" gorm:"column:damage_code;"`
	Damage               string `json:"damage" gorm:"column:damage;"`
	ZinspecVehicleTypeId int32  `json:"zinspec_vehicletype_id" gorm:"column:zinspec_vehicletype_id;"`
}

type DamageData struct {
	DamageRequest
}

type DamageDataResponse struct {
	DamageData
	Causes []CauseDataResponse `json:"causes"`
}

func (DamageData) TableName() string {
	return "zinspec_damage"
}
