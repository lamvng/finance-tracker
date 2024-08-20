package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GenericRepository[T any] interface {
	FindByID(id uuid.UUID) (*T, error)
	FindAll() ([]T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
}

type repository[T any] struct {
	db *gorm.DB
}
