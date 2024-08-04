package models

import "github.com/google/uuid"

// Type of assets
// Fixed by program. Eg: Commodities, real estate, bonds, stocks
// Note: This name is not "technically" (financially) correct in some cases. Anyways...
type AssetType struct {
	Base
	Name        string       `json:"name" gorm:"uniqueIndex"`
	Description *string      `json:"description"`
	AssetUnits  []*AssetUnit `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Asset unit
// Fixed by program. Eg: Ounce (gold), USD/CNY/EUR/JPY... (Currency)
type AssetUnit struct {
	Base
	Name              string              `json:"name" gorm:"uniqueIndex"`
	Description       *string             `json:"description"`
	AssetTypeID       uuid.UUID           `json:"assetTypeId"`
	AssetAccounts     []*AssetAccount     `gorm:"many2many:asset_account_portfolios"`
	AssetTransactions []*AssetTransaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
