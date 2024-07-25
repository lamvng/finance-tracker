package models

type User struct {
	Base
	FirstName                string                   `json:"firstname"`
	LastName                 string                   `json:"lastname"`
	Username                 string                   `json:"username"`
	Email                    string                   `json:"email" gorm:"uniqueIndex"`
	PasswordSalt             string                   `json:"passwordsalt"`
	PasswordHash             string                   `json:"passwordhash"`
	TransactionCategories    []TransactionCategory    `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TransactionSubCategories []TransactionSubCategory `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
