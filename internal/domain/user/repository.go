package user

import (
	"errors"
	"gotickts/internal/domain/user/dto"

	"gorm.io/gorm"
)

var ErrorAlreadyExist = errors.New("Email already exist")

type Repository interface {
	CreateUser(user *User) error
	LoginUser(email string) (*User, error)
	getUsers() ([]*dto.Response, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r repository) CreateUser(user *User) error {
	var result = r.db.Create(user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrorAlreadyExist
		}
		return result.Error
	}

	return nil
}

func (r repository) LoginUser(email string) (*User, error) {
	var user User
	result := r.db.Where(&User{Email: email}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r repository) getUsers() ([]*dto.Response, error) {
	var responses []*dto.Response

	// Model and Scan using DTO slice map
	err := r.db.Model(&User{}).
		Select("id", "name", "email", "created_at").
		Scan(&responses).Error

	if err != nil {
		return nil, err
	}

	return responses, nil
}
