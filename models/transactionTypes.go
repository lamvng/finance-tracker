package models

import (
	"github.com/google/uuid"
)

// Note: I name the object "Transaction" instead of "Transaction"
// To not confuse with DB Transaction.

// General type of an transaction
// Each type should be associated to an AccountType.
// Eg. Buying (groceries), selling (stocks)...
type TransactionType struct {
	Base
	Name                  string                 `json:"name" gorm:"uniqueIndex"`
	Description           *string                `json:"description"`
	AccountTypeID         uuid.UUID              `json:"accountTypeId" gorm:"index"`
	TransactionCategories []*TransactionCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

// Transaction category
// Eg. Healthcare, shopping, holidays
// Defined by user
type TransactionCategory struct {
	Base
	Name                     string                    `json:"name" gorm:"index"`
	Description              *string                   `json:"description"`
	UserID                   uuid.UUID                 `json:"userId" gorm:"index"`
	TransactionTypeID        uuid.UUID                 `json:"transactionTypeId" gorm:"index"`
	TransactionSubCategories []*TransactionSubCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Subcategory
// Detailed action, eg: Groceries
type TransactionSubCategory struct {
	Base
	Name                       string                       `json:"name" gorm:"index"`
	Description                *string                      `json:"description"`
	TransactionCategoryID      uuid.UUID                    `json:"transactionCategoryId" gorm:"index"`
	InvestmentTransactions     []*InvestmentTransaction     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	LiquidSpendingTransactions []*LiquidSpendingTransaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	InternalLiquidTransactions []*InternalLiquidTransaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	BudgetCategories           []*BudgetCategory            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
