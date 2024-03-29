package domain

type PlanResponse struct {
	Data        any    `json:"data"`
	Message     string `json:"message"`
	ElapsedTime string `json:"elapsed_time"`
}

type PlanRequest struct {
	CompanyCode      string `json:"company_code" form:"company_code" gorm:"column:company_code;"`
	Ba               string `json:"ba" form:"ba" gorm:"column:ba;"`
	RunningAccount   string `json:"running_account" form:"running_account" gorm:"column:running_account;"`
	Status           string `json:"status" form:"status" gorm:"column:status;"`
	PlanDate         string `json:"plan_date" form:"plan_date" gorm:"column:plan_date;"`
	InspectDate      string `json:"inspect_date" form:"inspect_date" gorm:"column:inspect_date;"`
	InspectTime      string `json:"inspect_time" form:"inspect_time" gorm:"column:inspect_time;"`
	ZtuAgriRunAcctId int32  `json:"ztuagri_runacct_id" form:"ztuagri_runacct_id" gorm:"column:ztuagri_runacct_id;"`
	Week             int32  `json:"week" form:"week" gorm:"column:week;"`
	SyncDate         string `json:"sync_date" form:"sync_date" gorm:"column:sync_date;"`
}

type PlanData struct {
	Id int32 `json:"id" gorm:"column:id"`
	PlanRequest
}

func (PlanRequest) TableName() string {
	return "zinspec_plan"
}
