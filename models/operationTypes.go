package models

import (
	"github.com/google/uuid"
)

// Note: I name the object "Operation" instead of "Transaction"
// To not confuse with DB Transaction.

// Type of assets
// Fixed by program. Eg: Liquid, commodities, real estate, bonds, stocks
// Note: This name is not "technically" (financially) correct in some cases. Anyways...
type AssetType struct {
	Base
	Name        string `json:"name" gorm:"uniqueIndex"`
	Description string `json:"description"`
}

// General type of an operation
// Each type should be associated to an AccountType.
// Eg. Buying (groceries), selling (stocks)...
type OperationGeneralType struct {
	Base
	Name          string              `json:"name" gorm:"uniqueIndex"`
	Description   string              `json:"description"`
	AccountTypeID uuid.UUID           `json:"accountTypeId"`
	Categories    []OperationCategory `gorm:"foreignKey:TypeID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Operation category
// Eg. Healthcare, shopping, holidays
// Defined by user
type OperationCategory struct {
	Base
	Name          string                 `json:"name" gorm:"index"`
	Description   string                 `json:"description"`
	User          User                   `json:"user"`
	TypeID        uuid.UUID              `json:"typeId"`
	SubCategories []OperationSubCategory `gorm:"foreignKey:CategoryID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Subcategory
// Detailed action, eg: Groceries
type OperationSubCategory struct {
	Base
	Name        string    `json:"name" gorm:"index"`
	Description string    `json:"description"`
	CategoryID  uuid.UUID `json:"categoryId"`
}
