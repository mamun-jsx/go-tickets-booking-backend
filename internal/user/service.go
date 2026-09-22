package user

import (
	"errors"

	"github.com/mamun-jsx/go-tickets-booking-backend.git/internal/user/dto"
)

type service struct {
	repo Repository
}

var ErrorInvalidEmailPassword = errors.New("Invalid email or password")

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

// create user service
func (s *service) CreateUser(req dto.CreateRequest) (*dto.Response, error) {
	// * take the repo create user function and call it and return the error
	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	err := user.hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// * save the user to database
	err = s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	response := dto.Response{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
	}
	return &response, nil
}

func (s *service) LoginUser(req dto.LoginRequest) (*dto.Response, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrorInvalidEmailPassword
	}
	//compare password || check password

	err = user.checkPassword(req.Password)
	if err != nil {
		return nil, ErrorInvalidEmailPassword
	}
	response := dto.Response{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
	}
	return &response, nil
}
