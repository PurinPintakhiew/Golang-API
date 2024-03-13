package models

type Users struct {
	Id       uint   `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
	UserId   []byte `json:"UserId"`
}

type RegistrationData struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}