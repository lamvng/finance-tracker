package database

import (
	"fmt"
	"lamvng/finance-tracker/configs"
	"lamvng/finance-tracker/model"
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
		&model.AccountType{},
		&model.AssetType{},
		&model.AssetUnit{},
		&model.AssetUnitExchangeRate{},
		&model.User{},
		&model.Lender{},
		&model.Account{},
		&model.TransactionType{},
		&model.TransactionCategory{},
		&model.TransactionSubCategory{},
		&model.InvestmentAccount{},
		&model.LiquidSavingAccount{},
		&model.LiquidSpendingAccount{},
		&model.LiquidSpendingBudget{},
		&model.InvestmentAccountPortfolio{},
		&model.InvestmentTransaction{},
		&model.LiquidSpendingTransaction{},
		&model.LendingTransaction{},
		&model.BudgetCategory{},
	)

	return DB
}
