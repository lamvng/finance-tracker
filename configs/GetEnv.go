package configs

import (
	"os"

	"github.com/golang/glog"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		glog.Fatal(err)
	}
}

func GetEnvVariables(key string) string {
	return os.Getenv(key)
}
