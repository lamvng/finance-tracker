package database

import (
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/redis/go-redis/v9"
)

var Cache *redis.Client

func InitRedisConnection() {
	redisUser := os.Getenv("REDIS_USER")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB := os.Getenv("REDIS_DB")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	redisURL := fmt.Sprintf("redis://%s:%s@%s:%s/%s", redisUser, redisPassword, redisHost, redisPort, redisDB)

	Opt, err := redis.ParseURL(redisURL)
	if err != nil {
		glog.Fatalf("Failed to connect to Redis: %s\n", err)
	}

	Cache = redis.NewClient(Opt)
}
