package models

type User struct {
	Base
	FirstName             string                `json:"firstName"`
	LastName              string                `json:"lastName"`
	Username              string                `json:"userName" gorm:"uniqueIndex"`
	Email                 string                `json:"email"`
	PasswordHash          string                `json:"passwordHash"`
	AssetAccounts         []*AssetAccount       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LiquidAccounts        []*LiquidAccount      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TransactionCategories []TransactionCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CreateUserInput struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Username  string `json:"userName" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
