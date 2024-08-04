package main

import (
	"lamvng/finance-tracker/db"
	"lamvng/finance-tracker/models"
)

func main() {
	db := db.Init()

	db.AutoMigrate(
		&models.AssetType{},
		&models.AssetUnit{},
		&models.LiquidCurrency{},
		&models.AccountType{},
		&models.User{},
		&models.AssetAccountPortfolio{},
		&models.AssetAccount{},
		&models.LiquidAccount{},
		&models.TransactionType{},
		&models.TransactionCategory{},
		&models.TransactionSubCategory{},
		&models.LiquidTransaction{},
		&models.AssetTransaction{},
	)
}
