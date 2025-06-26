package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Username  string `json:"username" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"` //TODO: add email validation
	Password  string `json:"-" gorm:"not null"`     //why "-" instead of "password" inside the json tag
}

//TODO: after finishing - get rid of GORM completely and use raw SQL or sqlc/sqlx, add email validation
