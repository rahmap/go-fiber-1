package user

import (
	"errors"
	"strings"
)

var ErrInvalidUserID = errors.New("user id must be greater than 0")
var ErrNameRequired = errors.New("name is required")

type Service interface {
	GetAll() ([]Response, error)
	GetByID(id uint) (*Response, error)
	Create(request CreateRequest) (*Response, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAll() ([]Response, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return toResponses(users), nil
}

func (s *service) GetByID(id uint) (*Response, error) {
	if id == 0 {
		return nil, ErrInvalidUserID
	}

	user, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := toResponse(*user)
	return &response, nil
}

func (s *service) Create(request CreateRequest) (*Response, error) {
	if strings.TrimSpace(request.Name) == "" {
		return nil, ErrNameRequired
	}

	user := User{
		Name:    strings.TrimSpace(request.Name),
		Age:     request.Age,
		Address: strings.TrimSpace(request.Address),
	}

	if err := s.repository.Create(&user); err != nil {
		return nil, err
	}

	response := toResponse(user)
	return &response, nil
}
