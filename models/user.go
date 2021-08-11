package models

type User struct {
	Id       uint   `json:"id"`
	FirtName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}
