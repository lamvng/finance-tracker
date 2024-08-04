package models

import (
	"time"

	"github.com/google/uuid"
)

// Liquid transactions details
type AssetTransaction struct {
	Base
	Name                     string    `json:"name" gorm:"uniqueIndex"`
	Description              *string   `json:"description"`
	AssetAmount              float32   `json:"assetAmount"`
	LiquidAmount             float32   `json:"liquidAmount"`
	ExecutedAt               time.Time `json:"executedAt" gorm:"index"`
	AssetUnitID              uuid.UUID `json:"assetUnitId"`
	AssetAccountID           uuid.UUID `json:"assetAccountId" gorm:"index"`
	TransactionSubCategoryID uuid.UUID `json:"transactionSubbCategoryId" gorm:"index"`
}

// Everyday transaction details
// Spending & Income
type LiquidTransaction struct {
	Base
	AccountID                uuid.UUID `json:"accountId" gorm:"index"`
	Description              *string   `json:"description"`
	ExecutedAt               time.Time `json:"executedAt" gorm:"index"`
	Amount                   float32   `json:"amount"`
	LiquidAccountID          uuid.UUID `json:"liquidAccountId" gorm:"index"`
	TransactionSubCategoryID uuid.UUID `json:"transactionSubbCategoryId" gorm:"index"`
}
