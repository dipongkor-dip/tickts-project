package user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required" gorm:"type:VARCHAR(100); NOT NULL"`
	Email    string `json:"email" validate:"required,email" gorm:"type:VARCHAR(200); UNIQUEINDEX; NOT NULL"`
	Password string `json:"password" validate:"required,min=6" gorm:"type:VARCHAR(100); NOT NULL"`
}

func (u *User) hashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

func (u *User) comparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
