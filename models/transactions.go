package models

import (
	"time"

	"github.com/google/uuid"
)

// Internal liquid transactions between accounts
type InternalLiquidTransaction struct {
	Base
	Name                     string    `json:"name" gorm:"index"`
	Description              *string   `json:"description"`
	ExecutedAt               time.Time `json:"executedAt" gorm:"index"`
	Amount                   float32   `json:"amount"`
	TransactionSubCategoryID uuid.UUID `json:"transactionSubCategoryID"`
	SrcAccountID             uuid.UUID `json:"SrcAccountId" gorm:"index"`
	DestAccountID            uuid.UUID `json:"DestAccountId" gorm:"index"`
}

type InvestmentTransaction struct {
	Name                     string    `json:"name" gorm:"index"`
	Description              *string   `json:"description"`
	ExecutedAt               time.Time `json:"executedAt" gorm:"index"`
	AssetAmount              float32   `json:"assetAmount"`
	LiquidAmount             float32   `json:"liquidAmount"`
	AssetUnitID              uuid.UUID `json:"assetUnitID" gorm:"index"`
	AccountID                uuid.UUID `json:"accountID" gorm:"index"`
	TransactionSubCategoryID uuid.UUID `json:"transactionSubCategoryID" gorm:"index"`
}

type LiquidSpendingTransaction struct {
	Name                     string    `json:"name" gorm:"index"`
	Description              *string   `json:"description"`
	ExecutedAt               time.Time `json:"executedAt" gorm:"index"`
	Amount                   float32   `json:"amount"`
	AccountID                uuid.UUID `json:"accountID" gorm:"index"`
	TransactionSubCategoryID uuid.UUID `json:"transactionSubCategoryID" gorm:"index"`
}

type LendingTransaction struct {
	Name        string    `json:"name" gorm:"index"`
	Description *string   `json:"description"`
	ExecutedAt  time.Time `json:"executedAt" gorm:"index"`
	DueAt       time.Time `json:"dueAt" gorm:"index"`
	Amount      float32   `json:"amount"`
	AccountID   uuid.UUID `json:"accountId" gorm:"index"`
	LenderID    uuid.UUID `json:"lenderId" gorm:"index"`
}
