package model

type User struct {
	Base
	FirstName             string                  `json:"firstName"`
	LastName              string                  `json:"lastName"`
	Username              string                  `json:"userName" gorm:"uniqueIndex"`
	Email                 string                  `json:"email"`
	PasswordHash          string                  `json:"-"`
	Accounts              []*Account              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TransactionCategories []*TransactionCategory  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LiquidSpendingBudget  []*LiquidSpendingBudget `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Lenders               []*Lender               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
