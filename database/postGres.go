package database

import (
	"fmt"
	"lamvng/finance-tracker/configs"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostGresConnection() *gorm.DB {
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

	return DB
}
