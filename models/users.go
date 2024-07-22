package models

type User struct {
	Base
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email" gorm:"uniqueIndex"`
	PasswordSalt string `json:"passwordsalt"`
	PasswordHash string `json:"passwordhash"`
}
