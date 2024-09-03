package configs

import (
	"os"

	"github.com/golang/glog"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		glog.Fatal(err)
	}
}

func GetEnvVariables(key string) string {
	return os.Getenv(key)
}
