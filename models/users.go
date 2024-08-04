package models

type User struct {
	Base
	FirstName             string                `json:"firstname"`
	LastName              string                `json:"lastname"`
	Username              string                `json:"username" gorm:"uniqueIndex"`
	Email                 string                `json:"email" gorm:"uniqueIndex"`
	PasswordSalt          string                `json:"passwordsalt"`
	PasswordHash          string                `json:"passwordhash"`
	AssetAccounts         []*AssetAccount       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LiquidAccounts        []*LiquidAccount      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TransactionCategories []TransactionCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
