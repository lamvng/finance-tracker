package repository

import (
	"lamvng/finance-tracker/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(id uuid.UUID) (model.User, error)
	Create(user model.User) error
	Update(id uuid.UUID, user model.User) (model.User, error)
	Delete(id uuid.UUID) error
}

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

func (r *UserRepositoryImpl) GetByID(id uuid.UUID) (model.User, error) {
	var user model.User
	if err := r.Db.First(&user, id).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Create(user model.User) error {
	err := r.Db.Create(&user).Error
	return err
}

func (r *UserRepositoryImpl) Update(id uuid.UUID, user model.User) (model.User, error) {

	existingUser, err := r.GetByID(id)

	// User does not exist
	if err != nil {
		return model.User{}, err
	}

	// Update user information
	if user.PasswordHash != "" {
		existingUser.PasswordHash = user.PasswordHash
	}
	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.Username = user.Username
	existingUser.Email = user.Email

	if err := r.Db.Model(&user).Updates(existingUser).Error; err != nil {
		return model.User{}, err
	}
	return existingUser, nil
}

func (r *UserRepositoryImpl) Delete(id uuid.UUID) error {
	existingUser, err := r.GetByID(id)
	if err != nil {
		return err
	}
	r.Db.Delete(&existingUser)
	return nil
}
