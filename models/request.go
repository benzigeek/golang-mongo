package models

type IRequestRegister struct {
	Data IRequestRegisterData
}

type IRequestRegisterData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
