package repositories

import (
	"fmt"
	"lamvng/finance-tracker/configs"
	"log"

	"github.com/redis/go-redis/v9"
)

func initRedisConnection() *redis.Client {
	redisUser := configs.GetEnvVariables("REDIS_USER")
	redisPassword := configs.GetEnvVariables("REDIS_PASSWORD")
	redisDB := configs.GetEnvVariables("REDIS_DB")
	redisHost := configs.GetEnvVariables("REDIS_HOST")
	redisPort := configs.GetEnvVariables("REDIS_PORT")

	redisURL := fmt.Sprintf("redis://%s:%s@%s:%s/%s", redisUser, redisPassword, redisHost, redisPort, redisDB)

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %s\n", err)
	}

	return redis.NewClient(opt)
}
