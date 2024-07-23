package models

import (
	"time"

	"github.com/google/uuid"
)

// Everyday transaction details
// Spending & Income
type EverydayTransaction struct {
	Base
	User        User                   `json:"user" gorm:"index"`
	AccountID   uuid.UUID              `json:"accountId" gorm:"index"`
	SubCategory TransactionSubCategory `json:"subcategory" gorm:"index"`
	ExecutedAt  time.Time              `json:"executedat" gorm:"index"`
	Amount      float32                `json:"amount" gorm:"index"`
	Description *string                `json:"description"`
}

// // Investment transaction details
// // Buying & Selling stocks, commodities...
// type InvestmentTransaction struct {
// 	Base
// 	User        User                 `json:"user" gorm:"index"`
// 	SubCategory TransactionSubCategory `json:"subcategory" gorm:"index"`
// 	DateTime    time.Time            `json:"time" gorm:"index"`
// 	Amount      float32              `json:"amount" gorm:"index"`
// }
