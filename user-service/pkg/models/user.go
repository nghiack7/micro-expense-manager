package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string `json:"user_name"`
	HashPassword string `json:"hash_password"`
	Email        string `json:"email"`
	NumberPhone  string `json:"number_phone"`
	Credit       int64  `json:"credit"`
}
