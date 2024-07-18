package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariables(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln(err)
	}

	return os.Getenv(key)
}
