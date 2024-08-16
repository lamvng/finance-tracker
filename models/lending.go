package models

import "github.com/google/uuid"

type Lender struct {
	Base
	Name                string                `json:"name" gorm:"index"`
	Description         *string               `json:"description"`
	UserID              uuid.UUID             `json:"userId" gorm:"index"`
	LendingTransactions []*LendingTransaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
