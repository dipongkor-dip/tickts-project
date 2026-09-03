package user

import "gotickts/internal/user/dto"

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r: r}
}

func (s *service) CreateUser(req dto.CreateRequest) (*dto.Response, error) {
	user := User{Name: req.Name, Email: req.Email, Password: req.Password}

	err := s.r.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	response := dto.Response{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.String(),
	}

	return &response, nil
}
