package model

import (
	"time"

	"github.com/google/uuid"
)

// Type of assets
// Fixed by program. Eg: Liquid, commodities, real estate, bonds, stocks
// Note: This name is not "technically" (financially) correct in some cases. Anyways...
type AssetType struct {
	Base
	Name        string       `json:"name" gorm:"uniqueIndex"`
	Description *string      `json:"description"`
	AssetUnits  []*AssetUnit `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

// Asset unit
// Fixed by program. Eg: Ounce (gold), USD/CNY/EUR/JPY... (Currency)
type AssetUnit struct {
	Base
	Name                        string                        `json:"name" gorm:"uniqueIndex"`
	Description                 *string                       `json:"description"`
	AssetTypeID                 uuid.UUID                     `json:"assetTypeId"`
	Accounts                    []*Account                    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:PrimaryAssetUnitID"`
	AssetUnitExchangeRateSrcs   []*AssetUnitExchangeRate      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:SrcAssetUnitID"`
	AssetUnitExchangeRateDests  []*AssetUnitExchangeRate      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;foreignKey:DestAssetUnitID"`
	InvestmentAccountPortfolios []*InvestmentAccountPortfolio `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	InvestmentTransactions      []*InvestmentTransaction      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	LiquidSpendingBudgets       []*LiquidSpendingBudget       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

// Exchange rate between assets
type AssetUnitExchangeRate struct {
	CreatedAt       time.Time  `json:"createdat"`
	UpdatedAt       time.Time  `json:"updatedat"`
	DeletedAt       *time.Time `json:"deletedat"`
	SrcAssetUnitID  uuid.UUID  `json:"srcAssetUnitID" gorm:"primary_key"`
	DestAssetUnitID uuid.UUID  `json:"destAssetUnitID" gorm:"primary_key"`
	Rate            float32    `json:"rate"`
}
