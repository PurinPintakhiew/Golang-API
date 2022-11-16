package models

type Users struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Status 	string `json:"status"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}