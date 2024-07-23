package models

import (
	"github.com/google/uuid"
)

// Note: I name the object "Transaction" instead of "Transaction"
// To not confuse with DB Transaction.

// General type of an transaction
// Each type should be associated to an AccountType.
// Eg. Buying (groceries), selling (stocks)...
type TransactionGeneralType struct {
	Base
	Name          string                `json:"name" gorm:"uniqueIndex"`
	Description   *string               `json:"description"`
	AccountTypeID uuid.UUID             `json:"accountTypeId"`
	Categories    []TransactionCategory `gorm:"foreignKey:TypeID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Transaction category
// Eg. Healthcare, shopping, holidays
// Defined by user
type TransactionCategory struct {
	Base
	Name          string                   `json:"name" gorm:"index"`
	Description   *string                  `json:"description"`
	User          User                     `json:"user"`
	TypeID        uuid.UUID                `json:"typeId"`
	SubCategories []TransactionSubCategory `gorm:"foreignKey:CategoryID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Subcategory
// Detailed action, eg: Groceries
type TransactionSubCategory struct {
	Base
	Name        string    `json:"name" gorm:"index"`
	Description *string   `json:"description"`
	CategoryID  uuid.UUID `json:"categoryId"`
}
