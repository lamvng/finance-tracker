package model

import (
	"time"

	"github.com/google/uuid"
)

// Type of user account
// Fixed by program. Each account should have a different logic.
// Everyday account, saving bonds, stock...
type AccountType struct {
	Base
	Name             string            `json:"name" gorm:"uniqueIndex"`
	Description      *string           `json:"description"`
	Accounts         []*Account        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	TransactionTypes []TransactionType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

// User financial accounts
// Defined by user
type Account struct {
	Base
	Name                          string                       `json:"name" gorm:"index"`
	Description                   *string                      `json:"description"`
	LiquidBalance                 float32                      `json:"liquidBalance"`
	AccountTypeID                 uuid.UUID                    `json:"accountTypeId" gorm:"index"`
	UserID                        uuid.UUID                    `json:"userId" gorm:"index"`
	PrimaryAssetUnitID            uuid.UUID                    `json:"primaryAssetUnitId" gorm:"index"`
	InternalLiquidTransactionsSrc []*InternalLiquidTransaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:SrcAccountID"`
	InternalLiquidTransactionDest []*InternalLiquidTransaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:DestAccountID"`
	LiquidSpendingAccount         *LiquidSpendingAccount       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	InvestmentAccount             *InvestmentAccount           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LiquidSavingAccount           *LiquidSavingAccount         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Investment accounts
type InvestmentAccount struct {
	AccountID                   uuid.UUID                     `gorm:"primaryKey;autoIncrement:false"`
	Account                     Account                       `gorm:"foreignKey:AccountID;references:ID"`
	InvestmentAccountPortfolios []*InvestmentAccountPortfolio `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	InvestmentTransactions      []*InvestmentTransaction      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AccountID"`
}

// Liquid spending accounts
type LiquidSpendingAccount struct {
	AccountID                  uuid.UUID                    `gorm:"primaryKey;autoIncrement:false"`
	Account                    Account                      `gorm:"foreignKey:AccountID;references:ID"`
	LiquidSpendingTransactions []*LiquidSpendingTransaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AccountID"`
	LendingTransactions        []*LendingTransaction        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AccountID"`
	LiquidSpendingBudgets      []*LiquidSpendingBudget      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Liquid saving accounts
type LiquidSavingAccount struct {
	AccountID uuid.UUID `gorm:"primaryKey;autoIncrement:false"`
	Account   Account   `gorm:"foreignKey:AccountID;references:ID"`
	APY       float32   `json:"apy"`
	StartAt   time.Time `json:"startAt"`
	EndAt     time.Time `json:"endAt"`
}
