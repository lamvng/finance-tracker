package model

import "github.com/google/uuid"

// User asset account portfolios
// Defined by user
type InvestmentAccountPortfolio struct {
	AssetUnitID         uuid.UUID `json:"assetUnitId" gorm:"primary_key"`
	InvestmentAccountID uuid.UUID `json:"investmentAccountId" gorm:"primary_key"`
	AssetAmount         float32   `json:"assetAmount"`
	LiquidAmount        float32   `json:"liquidAmount"`
}
