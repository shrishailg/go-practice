package dto

import "gorm.io/gorm"

type (
	User struct {
		gorm.Model
		Username     string `json:"name"`
		UserLastName string `json:"last_name"`
		Age          int    `json:"age"`
		Email        string `json:"email"`
	}
)
