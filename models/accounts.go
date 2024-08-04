package models

import "github.com/google/uuid"

// Type of user account
// Fixed by program. Each account should have a different logic.
// Everyday account, saving bonds, stock...
type AccountType struct {
	Base
	Name             string            `json:"name" gorm:"uniqueIndex"`
	Description      *string           `json:"description"`
	Accounts         []AssetAccount    `gorm:"foreignKey:AccountTypeID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TransactionTypes []TransactionType `gorm:"foreignKey:AccountTypeID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// User asset account portfolios
// Defined by user
type AssetAccountPortfolio struct {
	AssetUnitID    uuid.UUID `json:"assetUnitId" gorm:"primary_key"`
	AssetAccountID uuid.UUID `json:"assetAccountId" gorm:"primary_key"`
	AssetAmount    float32   `json:"assetAmount"`
	LiquidAmount   float32   `json:"liquidAmount"`
}

// User asset account
// Defined by user
type AssetAccount struct {
	Base
	Name             string             `json:"name" gorm:"uniqueIndex"`
	Description      *string            `json:"description"`
	LiquidBalance    float32            `json:"liquidBalance"`
	LiquidCurrencyID uuid.UUID          `json:"liquidCurrencyId" gorm:"index"`
	AccountTypeID    uuid.UUID          `json:"accountTypeId" gorm:"index"`
	UserID           uuid.UUID          `json:"userId" gorm:"index"`
	Transactions     []AssetTransaction `gorm:"foreignKey:AssetAccountID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type LiquidAccount struct {
	Base
	Name             string              `json:"name" gorm:"uniqueIndex"`
	Description      *string             `json:"description"`
	Balance          float32             `json:"balance"`
	LiquidCurrencyID uuid.UUID           `json:"liquidCurrencyId" gorm:"index"`
	AccountTypeID    uuid.UUID           `json:"accountTypeId" gorm:"index"`
	UserID           uuid.UUID           `json:"userId" gorm:"index"`
	Transactions     []LiquidTransaction `gorm:"foreignKey:LiquidAccountID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
