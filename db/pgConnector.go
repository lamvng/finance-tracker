package db

import (
	"fmt"
	"log"

	"lamvng/finance-tracker/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	postgresUser := configs.GetEnvVariables("POSTGRES_USER")
	postgresPassword := configs.GetEnvVariables("POSTGRES_PASSWORD")
	postgresDB := configs.GetEnvVariables("POSTGRES_DB")
	postgresHost := configs.GetEnvVariables("POSTGRES_HOST")
	postgresPort := configs.GetEnvVariables("POSTGRES_PORT")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", postgresUser, postgresPassword, postgresHost, postgresPort, postgresDB)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// db.AutoMigrate(&models.Book{})

	return db
}
