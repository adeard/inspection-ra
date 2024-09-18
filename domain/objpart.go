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
	ObjPartCode          string `json:"objpart_code" gorm:"column:objpart_code;"`
	Catalog              string `json:"catalog" gorm:"column:catalog;"`
}

type ObjPartData struct {
	Id              int32  `json:"id" gorm:"column:id"`
	VehicleTypeCode string `json:"vehicle_type_code" gorm:"column:vehicle_type_code"`
	ObjPartRequest
	VehicleTypeData VehicleTypeData `json:"vehicle_type_data" gorm:"foreignKey:Id;references:ZinspecVehicleTypeId;"`
}

type ObjPartDataResponse struct {
	Id              int32  `json:"id" gorm:"column:id"`
	VehicleTypeCode string `json:"vehicle_type_code" gorm:"column:vehicle_type_code"`
	ObjPartRequest
	Damages []DamageDataResponse `json:"damages"`
}

func (ObjPartData) TableName() string {
	return "zinspec_objpart"
}
