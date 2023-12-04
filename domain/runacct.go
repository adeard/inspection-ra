package domain

type RunAcctResponse struct {
	Data        any    `json:"data"`
	Message     string `json:"message"`
	ElapsedTime string `json:"elapsed_time"`
}

type RunAcctRequest struct {
	RunningAccount string `json:"running_account" gorm:"column:running_account;" form:"running_account"`
	CompanyCode    string `json:"company_code" gorm:"column:company_code;" form:"company_code"`
	Estate         string `json:"estate" gorm:"column:estate;" form:"estate"`
	Anln1          int32  `json:"anln1" gorm:"column:anln1;" form:"anln1"`
	ClassRa        string `json:"class_ra" gorm:"column:class_ra" form:"class_ra"`
	StatusFlag     string `json:"status_flag" gorm:"column:status_flag" form:"status_flag"`
	Kostl          string `json:"kostl" gorm:"column:kostl" form:"kostl"`
	ClassGroup     string `json:"class_group" gorm:"column:class_group" form:"class_group"`
	LicensePlate   string `json:"license_plate" gorm:"column:license_plate" form:"license_plate"`
	RunningText    string `json:"running_text" gorm:"column:running_text" form:"running_text"`
	Equnr          string `json:"equnr" gorm:"column:equnr" form:"equnr"`
}

type RunAcctData struct {
	Id int32 `json:"id" gorm:"column:id"`
	RunAcctRequest
}

func (RunAcctData) TableName() string {
	return "ztuagri_runacct"
}
