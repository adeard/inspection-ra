package domain

type VehicleTypeResponse struct {
	Data        any    `json:"data"`
	Message     string `json:"message"`
	ElapsedTime string `json:"elapsed_time"`
}

type VehicleTypeRequest struct {
	VehicleType string `json:"vehicle_type" gorm:"column:vehicle_type;"`
	Code        string `json:"code" gorm:"column:code;"`
	ObjPart     string `json:"obj_part" gorm:"column:obj_part;"`
}

type VehicleTypeData struct {
	Id int32 `json:"id" gorm:"column:id;"`
	VehicleTypeRequest
}

func (VehicleTypeData) TableName() string {
	return "zinspec_vehicletype"
}
