package service

import (
	"errors"
	"lamvng/finance-tracker/data/request"
	"lamvng/finance-tracker/data/response"
	"lamvng/finance-tracker/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserService interface {
	GetByID(id uuid.UUID) (response.GetUserByIDResponse, error)
	Create(user request.CreateUserRequest)
	Update(user request.UpdateUserRequest)
	Delete(id uuid.UUID)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewTagServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) (userService UserService, err error) {
	if validate == nil {
		return nil, errors.New("validator instance cannot be nil")
	}
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}, err
}

func (s *UserServiceImpl) GetByID(id uuid.UUID) {
	user, err := s.UserRepository.GetByID(id)
	if err != nil {
		return response.GetUserByIDResponse{}, nil
	}
}
func (s *UserServiceImpl) Create(user request.CreateUserRequest) {
}
