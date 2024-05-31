package domain

type Auth struct {
	Domain   string `json:"domain"`
	Ldap     bool   `json:"ldap"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthRequest struct {
	Domain   string `json:"domain"`
	Ldap     bool   `json:"ldap"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthData struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

type AuthLoggedParsedData struct {
	AuthLoggedData
	Employees []AuthEmployeeData `json:"employees"`
}

type AuthLoggedData struct {
	Id       string            `json:"id"`
	Code     string            `json:"code"`
	Username string            `json:"username"`
	Email    string            `json:"email"`
	Fullname string            `json:"fullname"`
	Company  []AuthCompanyData `json:"company"`
}

type AuthCompanyData struct {
	Id          string           `json:"id"`
	Code        string           `json:"code"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Estates     []AuthEstateData `json:"estates"`
}

type AuthEstateData struct {
	Id          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AuthEmployeeData struct {
	Id          string `json:"id"`
	Code        string `json:"code"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Nik         string `json:"nik"`
	Nickname    string `json:"nickname"`
	Fullname    string `json:"fullname"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Department  string `json:"department"`
}

type AuthLoggedResponse struct {
	Data    AuthLoggedData `json:"data"`
	Message string         `json:"message"`
	Status  int            `json:"status"`
}

type AuthResponse struct {
	Data        any    `json:"data"`
	Message     string `json:"message"`
	ElapsedTime string `json:"elapsed_time"`
}

type NewUserResponse struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
	Datas   string `json:"datas"`
}

type NewDetailUserResponse struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
	Objek   any    `json:"objek"`
}
