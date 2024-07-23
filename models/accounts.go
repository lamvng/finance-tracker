package models

import "github.com/google/uuid"

// Type of user account
// Fixed by program. Each account should have a different logic.
// Everyday account, saving bonds, stock...
type AccountType struct {
	Base
	Name             string                   `json:"name" gorm:"uniqueIndex"`
	Description      *string                  `json:"description"`
	Accounts         []Account                `gorm:"foreignKey:TypeID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TransactionTypes []TransactionGeneralType `gorm:"foreignKey:AccountTypeID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// User account
// Defined by user
type Account struct {
	Base
	Name         string                `json:"name" gorm:"uniqueIndex"`
	Description  *string               `json:"description"`
	TypeID       uuid.UUID             `json:"typeId" gorm:"index"`
	AssetUnitID  uuid.UUID             `json:"assetUnitId" gorm:"index"`
	Balance      float32               `json:"balance"`
	Transactions []EverydayTransaction `gorm:"foreignKey:AccountID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
