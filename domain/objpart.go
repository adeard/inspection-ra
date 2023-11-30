package domain

type ObjPartResponse struct {
	Data        any    `json:"data"`
	Message     string `json:"message"`
	ElapsedTime string `json:"elapsed_time"`
}

type ObjPartRequest struct {
	ObjPart              string `json:"obj_part" gorm:"column:obj_part;"`
	VehicleType          string `json:"vehicle_type" gorm:"column:vehicle_type;"`
	ZinspecVehicleTypeId int32  `json:"zinspec_vehicletype_id" gorm:"column:zinspec_vehicletype_id;"`
}

type ObjPartData struct {
	Id int32 `json:"id" gorm:"column:id"`
	ObjPartRequest
}

func (ObjPartData) TableName() string {
	return "zinspec_objpart"
}
