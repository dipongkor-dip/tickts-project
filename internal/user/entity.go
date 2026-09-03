package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required" gorm:"type:VARCHAR(100); NOT NULL"`
	Email    string `json:"email" validate:"required,email" gorm:"type:VARCHAR(200); UNIQUEINDEX; NOT NULL"`
	Password string `json:"password" validate:"required,min=6" gorm:"type:VARCHAR(100); NOT NULL"`
}

// 74-8 - 5.00
