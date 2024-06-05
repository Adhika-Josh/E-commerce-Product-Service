package model

type CreateAdminRequest struct {
	Name     string `json:"name"`
	PhoneNo  string `json:"phone_no"`
	Dob      string `json:"dob"`
	Email    string `json:"emai"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
type CreateAdminResponse struct {
	Status   int    `json:"status"`
	Message  string `json:"message"`
	AdminPID string `json:"admin_pid"`
}

type LoginAdminRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginAdminResponse struct {
	Status   int    `json:"status"`
	Message  string `json:"message"`
	AdminPID string `json:"admin_pid"`
}
