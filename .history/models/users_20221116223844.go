package models

type Users struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Status 	string `json:"status"`
	Telephone string `json:"telephone"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}