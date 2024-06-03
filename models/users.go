package models

type Users struct {
    Id       uint   `json:"id" gorm:"primaryKey"`
    UserName string `json:"user_name"`
    Email    string `json:"email" gorm:"uniqueIndex"`
    Password []byte `json:"-"`
}

type RegistrationData struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}