package models

import (
	"gorm.io/gorm"
)

type TransactionGeneralType struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex"`
	Description string
	Categories  []TransactionCategory
}

type TransactionCategory struct {
	gorm.Model
	Name          string `gorm:"index"`
	Description   string
	SubCategories []TransactionSubCategory
}

type TransactionSubCategory struct {
	gorm.Model
	Name        string `gorm:"index"`
	Description string
}

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"uniqueIndex"`
}
