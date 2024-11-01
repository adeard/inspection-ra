package domain

type MobResponse struct {
	Data        any    `json:"data"`
	Message     string `json:"message"`
	ElapsedTime string `json:"elapsed_time"`
}

type MobRequest struct {
	NoInspec             string `json:"no_inspec" gorm:"column:no_inspec;"`
	Ba                   string `json:"ba" gorm:"column:ba;"`
	RunningAccount       string `json:"running_account" gorm:"column:running_account;"`
	VehicleType          string `json:"vehicle_type" gorm:"column:vehicle_type;"`
	ObjPart              string `json:"obj_part" gorm:"column:obj_part;"`
	DamDate              string `json:"dam_date" gorm:"column:dam_date;"`
	Status               int32  `json:"status" gorm:"column:status;"`
	Loc                  string `json:"loc" gorm:"column:loc;"`
	Photo                string `json:"photo" gorm:"column:photo"`
	Note                 string `json:"note" gorm:"column:note"`
	PlanDate             string `json:"plan_date" gorm:"column:plan_date"`
	CreatedBy            string `json:"created_by" gorm:"column:created_by"`
	CreatedDate          string `json:"created_date" gorm:"column:created_date"`
	CreatedTime          string `json:"created_time" gorm:"column:created_time"`
	ZtuAgriRunAcctId     int32  `json:"ztuagri_runacct_id" gorm:"column:ztuagri_runacct_id"`
	ZinspecVehicleTypeId int32  `json:"zinspec_vehicletype_id" gorm:"column:zinspec_vehicletype_id"`
	Coordinate           string `json:"coordinate" gorm:"column:coordinate"`
	IsGps                int32  `json:"is_gps" gorm:"column:is_gps"`
	DamDetail            string `json:"dam_detail" gorm:"column:dam_detail"`
	VersionApp           string `json:"version_app" gorm:"column:version_app"`
	SyncDate             string `json:"sync_date" form:"sync_date" gorm:"column:sync_date;"`
	NotifType            string `json:"notif_type" form:"notif_type" gorm:"column:notif_type;"`
	NotifDesc            string `json:"notif_desc" form:"notif_desc" gorm:"column:notif_desc;"`
}

type MobData struct {
	Id int32 `json:"id" gorm:"column:id;"`
	MobRequest
}

func (MobRequest) TableName() string {
	return "zinspec_mob"
}

type MobItemDamagedRequest struct {
	NoInspec   string `json:"no_inspec" gorm:"column:no_inspec;"`
	ObjPart    string `json:"obj_part" gorm:"column:obj_part;"`
	Damage     string `json:"damage" gorm:"column:damage;"`
	DamageText string `json:"damage_text" gorm:"column:damage_text;"`
	Causes     string `json:"causes" gorm:"column:causes;"`
	CausesText string `json:"causes_text" gorm:"column:causes_text;"`
	LineItem   int    `json:"line_item" gorm:"column:line_item;"`
}

type MobItemDamagedData struct {
	Id int32 `json:"id" gorm:"column:id;"`
	MobItemDamagedRequest
}

func (MobItemDamagedRequest) TableName() string {
	return "zinspec_mob_item_damaged"
}
