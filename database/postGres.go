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

	Db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %s\n", err)
	}

	Db.AutoMigrate(
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

	return Db
}
