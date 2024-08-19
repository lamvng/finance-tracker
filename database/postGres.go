package database

import (
	"fmt"
	"lamvng/finance-tracker/configs"
	"lamvng/finance-tracker/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initPostGresConnection() *gorm.DB {
	postgresUser := configs.GetEnvVariables("POSTGRES_USER")
	postgresPassword := configs.GetEnvVariables("POSTGRES_PASSWORD")
	postgresDB := configs.GetEnvVariables("POSTGRES_DB")
	postgresHost := configs.GetEnvVariables("POSTGRES_HOST")
	postgresPort := configs.GetEnvVariables("POSTGRES_PORT")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", postgresUser, postgresPassword, postgresHost, postgresPort, postgresDB)

	DB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %s\n", err)
	}

	DB.AutoMigrate(
		&models.AccountType{},
		&models.AssetType{},
		&models.AssetUnit{},
		&models.AssetUnitExchangeRate{},
		&models.User{},
		&models.Lender{},
		&models.Account{},
		&models.TransactionType{},
		&models.TransactionCategory{},
		&models.TransactionSubCategory{},
		&models.InvestmentAccount{},
		&models.LiquidSavingAccount{},
		&models.LiquidSpendingAccount{},
		&models.LiquidSpendingBudget{},
		&models.InvestmentAccountPortfolio{},
		&models.InvestmentTransaction{},
		&models.LiquidSpendingTransaction{},
		&models.LendingTransaction{},
		&models.BudgetCategory{},
	)

	return DB
}
