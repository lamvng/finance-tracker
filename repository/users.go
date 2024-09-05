package repository

import (
	"lamvng/finance-tracker/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetByID(id uuid.UUID) (model.User, error)
	GetByUsername(username string) (model.User, error)
	GetByEmail(email string) (model.User, error)
	Create(user model.User) error
	Update(user model.User) error
	Delete(user model.User) error
}

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{Db: Db}
}

func (r *UserRepository) GetByID(id uuid.UUID) (model.User, error) {
	var user model.User
	if err := r.Db.First(&user, id).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetByUsername(username string) (model.User, error) {
	var user model.User
	if err := r.Db.Where("username = ?", username).First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (model.User, error) {
	var user model.User
	if err := r.Db.Where("email = ?", email).First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Create(user model.User) error {
	err := r.Db.Create(&user).Error
	return err
}

func (r *UserRepository) Update(user model.User) error {
	var updatedUser model.User
	if user.PasswordHash != "" {
		updatedUser.PasswordHash = user.PasswordHash
	}
	updatedUser.FirstName = user.FirstName
	updatedUser.LastName = user.LastName
	updatedUser.Username = user.Username
	updatedUser.Email = user.Email

	err := r.Db.Model(&user).Updates(updatedUser).Error
	return err
}

func (r *UserRepository) Delete(user model.User) error {
	err := r.Db.Delete(&user).Error
	return err
}
