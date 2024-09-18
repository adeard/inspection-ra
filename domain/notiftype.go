package domain

type NotifTypeResponse struct {
	Data        any    `json:"data"`
	Message     string `json:"message"`
	ElapsedTime string `json:"elapsed_time"`
}

type NotifTypeRequest struct {
	Notif_Type          string `form:"notif_type" json:"notif_type" gorm:"column:Notif_Type;"`
	Plant               string `form:"plant" json:"plant" gorm:"column:Plant;"`
	Notif_Desc          string `form:"notif_desc" json:"notif_desc" gorm:"column:Notif_Desc;"`
	Application         string `form:"application" json:"application" gorm:"column:Application;"`
	Notification_Origin string `form:"notification_origin" json:"notification_origin" gorm:"column:Notification_Origin;"`
	Reference_Time      string `form:"reference_time" json:"reference_time" gorm:"column:Reference_Time;"`
	Catalog_Profile     string `form:"catalog_profile" json:"catalog_profile" gorm:"column:Catalog_Profile;"`
	NotifCategoryText   string `form:"notif_category_text" json:"notif_category_text" gorm:"column:NotifCategoryText;"`
	Update_Group        string `form:"update_group" json:"update_group" gorm:"column:Update_Group;"`
	Name                string `form:"name" json:"name" gorm:"column:Name;"`
	Description         string `form:"description" json:"description" gorm:"column:Description;"`
	Early_NoAssignment  string `form:"early_no_assignment" json:"early_no_assignment" gorm:"column:Early_NoAssignment;"`
	Number_Range        string `form:"number_range" json:"number_range" gorm:"column:Number_Range;"`
	PartnerDetermProced string `form:"partner_determ_proced" json:"partner_determ_proced" gorm:"column:PartnerDetermProced;"`
}

type NotifTypeData struct {
	NotifTypeRequest
}

type NotifTypeDataResponse struct {
	NotifTypeRequest
}

func (NotifTypeData) TableName() string {
	return "zinspec_notiftype"
}
