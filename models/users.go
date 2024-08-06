package models

type User struct {
	Base
	FirstName             string                `json:"firstname"`
	LastName              string                `json:"lastname"`
	Username              string                `json:"username" gorm:"uniqueIndex"`
	Email                 string                `json:"email"`
	PasswordHash          string                `json:"passwordhash"`
	AssetAccounts         []*AssetAccount       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LiquidAccounts        []*LiquidAccount      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TransactionCategories []TransactionCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CreateUserInput struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
