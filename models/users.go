package models

type User struct {
	Base
	FirstName             string                `json:"firstname"`
	LastName              string                `json:"lastname"`
	Username              string                `json:"username"`
	Email                 string                `json:"email" gorm:"uniqueIndex"`
	PasswordSalt          string                `json:"passwordsalt"`
	PasswordHash          string                `json:"passwordhash"`
	AssetAccounts         []LiquidAccount       `gorm:"foreignKey:assetAccountId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LiquidAccounts        []LiquidAccount       `gorm:"foreignKey:liquidAccountId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TransactionCategories []TransactionCategory `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
