package model

import "github.com/google/uuid"

type LiquidSpendingBudget struct {
	Base
	Name                    string    `json:"name" gorm:"index"`
	Description             *string   `json:"description"`
	Limit                   float32   `json:"limit"`
	UserID                  uuid.UUID `json:"userId" gorm:"index"`
	LiquidSpendingAccountID uuid.UUID `json:"liquidSpendingAccountId" gorm:"index"`
	AssetUnitID             uuid.UUID `json:"assetUnitID" gorm:"index"`
}

type BudgetCategory struct {
	LiquidSpendingBudgetID   uuid.UUID `json:"liquidSpendingBudgetId" gorm:"index"`
	TransactionSubCategoryID uuid.UUID `json:"transactionSubCategoryId" gorm:"index"`
}
