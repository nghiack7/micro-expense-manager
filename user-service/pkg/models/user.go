package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `json:"username"`
	HashPassword string `json:"hash_password"`
	Email        string `json:"email"`
	NumberPhone  string `json:"number_phone"`
	Credit       uint   `json:"credit"`
}
