package models

import "gorm.io/gorm"

//user model
type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"-" gorm:"not null"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
	IsActive  bool   `json:"is_active" gorm:"default:false"`
}

//auth model
type AuthModel struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

//login model
type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
