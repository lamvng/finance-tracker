package database

import (
	"fmt"
	"lamvng/finance-tracker/configs"

	"github.com/golang/glog"
	"github.com/redis/go-redis/v9"
)

func InitRedisConnection() *redis.Client {
	redisUser := configs.GetEnv("REDIS_USER")
	redisPassword := configs.GetEnv("REDIS_PASSWORD")
	redisDB := configs.GetEnv("REDIS_DB")
	redisHost := configs.GetEnv("REDIS_HOST")
	redisPort := configs.GetEnv("REDIS_PORT")

	redisURL := fmt.Sprintf("redis://%s:%s@%s:%s/%s", redisUser, redisPassword, redisHost, redisPort, redisDB)

	Opt, err := redis.ParseURL(redisURL)
	if err != nil {
		glog.Fatalf("Failed to connect to Redis: %s\n", err)
	}

	return redis.NewClient(Opt)
}
