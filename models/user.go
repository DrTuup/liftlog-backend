package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID        int
	Firstname string
	Lastname  string
	Email     string
	Username  string
	Password  string
}
