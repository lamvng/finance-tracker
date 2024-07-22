package models

import "time"

// Everyday operation details
// Spending & Income
type EverydayOperation struct {
	Base
	User        User                 `json:"user" gorm:"index"`
	SubCategory OperationSubCategory `json:"subcategory" gorm:"index"`
	ExecutedAt  time.Time            `json:"executedat" gorm:"index"`
	Amount      float32              `json:"amount" gorm:"index"`
}

// // Investment operation details
// // Buying & Selling stocks, commodities...
// type InvestmentOperation struct {
// 	Base
// 	User        User                 `json:"user" gorm:"index"`
// 	SubCategory OperationSubCategory `json:"subcategory" gorm:"index"`
// 	DateTime    time.Time            `json:"time" gorm:"index"`
// 	Amount      float32              `json:"amount" gorm:"index"`
// }
