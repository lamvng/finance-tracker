package service

import (
	"errors"
	"lamvng/finance-tracker/data/request"
	"lamvng/finance-tracker/data/response"
	"lamvng/finance-tracker/model"
	"lamvng/finance-tracker/repository"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	Auth(request.AuthenticationRequest) (response.AuthenticationResponse, error)
	GetByID(id uuid.UUID) (response.GetUserResponse, error)
	Create(userReq request.CreateUserRequest) error
	Update(userReq request.UpdateUserRequest, userId uuid.UUID) error
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

func (s *UserService) Auth(authReq request.AuthenticationRequest) (response.AuthenticationResponse, error) {

	// Find existing username
	user, err := s.UserRepository.GetByUsername(authReq.Username)
	if err != nil {
		return response.AuthenticationResponse{}, err
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(authReq.Password)); err != nil {
		return response.AuthenticationResponse{}, errors.New("password not correct")
	}

	// Calculate and return JWT token
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})
	token, err := generateToken.SignedString([]byte(os.Getenv("JWT_TOKEN_SECRET")))
	if err != nil {
		return response.AuthenticationResponse{}, err
	}
	authResponse := response.AuthenticationResponse{
		Token: token,
	}
	return authResponse, nil
}

func (s *UserService) GetByID(id uuid.UUID) (response.GetUserResponse, error) {

	user, err := s.UserRepository.GetByID(id)
	if err != nil {
		return response.GetUserResponse{}, err
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
	_, err := s.UserRepository.GetByEmail(userReq.Email)
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
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.PasswordHash = string(passwordHash)

	newUser.FirstName = userReq.FirstName
	newUser.LastName = userReq.LastName
	newUser.Username = userReq.Username
	newUser.Email = userReq.Email

	err = s.UserRepository.Create(newUser)
	return err
}

func (s *UserService) Update(userReq request.UpdateUserRequest, userId uuid.UUID) error {
	user, err := s.UserRepository.GetByID(userId)
	if err != nil {
		return err
	}

	user.FirstName = userReq.FirstName
	user.LastName = userReq.LastName
	user.Email = userReq.Email
	if userReq.Password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.PasswordHash = string(passwordHash)
	}

	// Update user in database
	err = s.UserRepository.Update(user)
	return err
}
