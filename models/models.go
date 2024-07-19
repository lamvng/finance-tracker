package models

import (
	"gorm.io/gorm"
)

// Note: I name the object "Operation" instead of "Transaction"
// To not confuse with DB Transaction.

// Type of assets
// Fixed by program. Eg: Liquid, commodities, real estate, bonds, stocks
type AssetType struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex"`
	Description string
}

// Type of user accounts
// Eg. Everyday account, saving bonds, stock...
type AccountType struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex"`
	Description string
}

// User account
type Account struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex"`
	Description string
	User        User
	Type        AccountType
}

// General type of an operation
// Eg. Buying (groceries), selling (stocks)...
type OperationGeneralType struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex"`
	Description string
	Categories  []OperationCategory
	AccountType AccountType
}

// Operation category
// Eg. Healthcare, shopping, holidays
type OperationCategory struct {
	gorm.Model
	Name          string `gorm:"index"`
	Description   string
	SubCategories []OperationSubCategory
}

// Subcategory
// Detailed action, eg: groceries
type OperationSubCategory struct {
	gorm.Model
	Name        string `gorm:"index"`
	Description string
}

// Operation detail
type Operation struct {
	gorm.Model
	Account     Account
	SubCategory OperationSubCategory
}

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"uniqueIndex"`
}
