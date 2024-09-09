package database

import (
	"fmt"
	"lamvng/finance-tracker/configs"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostGresConnection() *gorm.DB {
	postgresUser := configs.GetEnv("POSTGRES_USER")
	postgresPassword := configs.GetEnv("POSTGRES_PASSWORD")
	postgresDB := configs.GetEnv("POSTGRES_DB")
	postgresHost := configs.GetEnv("POSTGRES_HOST")
	postgresPort := configs.GetEnv("POSTGRES_PORT")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", postgresUser, postgresPassword, postgresHost, postgresPort, postgresDB)

	DB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %s\n", err)
	}

	return DB
}
