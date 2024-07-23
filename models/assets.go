package models

import "github.com/google/uuid"

// Type of assets
// Fixed by program. Eg: Liquid, commodities, real estate, bonds, stocks
// Note: This name is not "technically" (financially) correct in some cases. Anyways...
type AssetType struct {
	Base
	Name        string      `json:"name" gorm:"uniqueIndex"`
	Description string      `json:"description"`
	Units       []AssetUnit `gorm:"foreignKey:TypeID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Asset unit
// Fixed by program. Eg: Ounce (gold), USD/CNY/EUR/JPY... (Currency)
type AssetUnit struct {
	Base
	Name        string    `json:"name" gorm:"uniqueIndex"`
	Description string    `json:"description"`
	TypeID      uuid.UUID `json:"typeId"`
	Accounts    []Account `gorm:"foreignKey:AssetUnitID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
