package models

type LiquidCurrency struct {
	Base
	Name           string           `json:"name" gorm:"uniqueIndex"`
	Description    *string          `json:"description"`
	LiquidAccounts []*LiquidAccount `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AssetAccounts  []*AssetAccount  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
