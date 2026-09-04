package user

import (
	"fmt"
	"gotickts/internal/auth"
	"gotickts/internal/user/dto"
)

var ErrInvalidCredentials = fmt.Errorf("Invalid email or password")

type service struct {
	repo       Repository
	jwtService auth.JWTService
}

func NewService(r Repository, jwtService auth.JWTService) *service {
	return &service{repo: r, jwtService: jwtService}
}

func (s *service) CreateUser(req dto.CreateRequest) (*dto.RegisterRes, error) {
	user := User{Name: req.Name, Email: req.Email}

	err := user.hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	response := dto.RegisterRes{
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
	}

	return &response, nil
}

func (s *service) LoginUser(req dto.LoginRequest) (*dto.LoginRes, error) {
	user, err := s.repo.LoginUser(req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrInvalidCredentials
	}

	err = user.comparePassword(req.Password)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// generate token
	token, err := s.jwtService.GenerateToken(user.ID, user.Email, user.Name)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate token: %w", err)
	}

	res := dto.LoginRes{
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return &res, nil
}
