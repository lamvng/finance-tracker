package model

import (
	"time"

	"github.com/google/uuid"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdat"`
	UpdatedAt time.Time  `json:"updatedat"`
	DeletedAt *time.Time `json:"deletedat" gorm:"index"`
}
