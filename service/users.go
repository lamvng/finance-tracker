package service

import (
	"errors"
	"lamvng/finance-tracker/data/request"
	"lamvng/finance-tracker/data/response"
	"lamvng/finance-tracker/model"
	"lamvng/finance-tracker/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetByID(id uuid.UUID) (response.GetUserResponse, error)
	Create(user request.CreateUserRequest) error
	// Update(user request.UpdateUserRequest) error
	// Delete(id uuid.UUID) error
}

type UserService struct {
	UserRepository repository.UserRepositoryInterface
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepositoryInterface, validate *validator.Validate) (userService UserServiceInterface) {
	return &UserService{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (s *UserService) GetByID(id uuid.UUID) (response.GetUserResponse, error) {
	user, err := s.UserRepository.GetByID(id)
	if err != nil {
		return response.GetUserResponse{}, nil
	}
	res := response.GetUserResponse{
		ID:        uuid.UUID.String(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
	}
	return res, nil
}

func (s *UserService) Create(userReq request.CreateUserRequest) error {

	// Verify if email exists
	_, err := s.UserRepository.GetByEmail(userReq.Username)
	if err == nil {
		return errors.New("email already in use")
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}

	// Verify if username exists
	_, err = s.UserRepository.GetByUsername(userReq.Username)
	if err == nil {
		return errors.New("username already in use")
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}

	var newUser model.User

	// Create Password Hash
	if userReq.Password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		newUser.PasswordHash = string(passwordHash)
	}

	newUser.FirstName = userReq.FirstName
	newUser.LastName = userReq.LastName
	newUser.Username = userReq.Username
	newUser.Email = userReq.Email

	err = s.UserRepository.Create(newUser)
	return err
}

// func (s *UserService) Update(user request.CreateUserRequest) error {

// }
