package models

type User struct {
	Base
	FirstName             string                  `json:"firstName"`
	LastName              string                  `json:"lastName"`
	Username              string                  `json:"userName" gorm:"uniqueIndex"`
	Email                 string                  `json:"email"`
	PasswordHash          string                  `json:"passwordHash"`
	Accounts              []*Account              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TransactionCategories []*TransactionCategory  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LiquidSpendingBudget  []*LiquidSpendingBudget `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Lenders               []*Lender               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CreateUserInput struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Username  string `json:"userName" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
