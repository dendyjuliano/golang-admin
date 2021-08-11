package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       uint   `json:"id"`
	FirtName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
	RoleId   uint   `json:"role_id"`
	Role     Role   `json:"role" gorm:"foreignKey:RoleId"`
}

func (user *User) SetPassword(password string) {
	Hashedpassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = Hashedpassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
