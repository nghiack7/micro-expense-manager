package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `json:"user_name"`
	HashPassword string `json:"hash_password"`
	Email        string `json:"email"`
	NumberPhone  int32  `json:"number_phone"`
	Credit       uint   `json:"credit"`
}
