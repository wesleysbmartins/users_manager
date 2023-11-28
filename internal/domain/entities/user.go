package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
