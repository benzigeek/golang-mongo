package models

type IRequestRegister struct {
	Data IRequestRegisterData
}

type IRequestRegisterData struct {
	Email    string
	Password string
}
