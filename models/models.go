package models

import (
	"time"

	"github.com/google/uuid"
)

// Note: I name the object "Operation" instead of "Transaction"
// To not confuse with DB Transaction.

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"createdat"`
	UpdatedAt time.Time  `json:"updatedat"`
	DeletedAt *time.Time `json:"deletedat" sql:"index"`
}

// Type of assets
// Fixed by program. Eg: Liquid, commodities, real estate, bonds, stocks
// Note: This name is not "technically" (financially) correct in some cases. Anyways...
type AssetType struct {
	Base
	Name        string `json:"name" gorm:"uniqueIndex"`
	Description string `json:"description"`
}

// Type of user wallet
// Fixed by program. Each account should have a different logic.
// Everyday wallet, saving bonds, stock...
type WalletType struct {
	Base
	Name        string `json:"name" gorm:"uniqueIndex"`
	Description string `json:"description"`
}

// General type of an operation
// Each type should be associated to an WalletType.
// Eg. Buying (groceries), selling (stocks)...
type OperationGeneralType struct {
	Base
	Name        string     `json:"name" gorm:"uniqueIndex"`
	Description string     `json:"description"`
	WalletType  WalletType `json:"wallettype"`
}

// Operation category
// Eg. Healthcare, shopping, holidays
// Defined by user
type OperationCategory struct {
	Base
	Name        string               `json:"name" gorm:"index"`
	Description string               `json:"description"`
	User        User                 `json:"user"`
	Type        OperationGeneralType `json:"type"`
}

// Subcategory
// Detailed action, eg: Groceries
type OperationSubCategory struct {
	Base
	Name        string            `json:"name" gorm:"index"`
	Description string            `json:"description"`
	Category    OperationCategory `json:"category"`
}

// Operation detail
type Operation struct {
	Base
	User        User                 `json:"user"`
	SubCategory OperationSubCategory `json:"subcategory"`
}

type User struct {
	Base
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email" gorm:"uniqueIndex"`
	PasswordSalt string `json:"passwordsalt"`
	PasswordHash string `json:"passwordhash"`
}
