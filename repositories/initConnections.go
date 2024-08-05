package repositories

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var db *gorm.DB
var rdb *redis.Client

func init() {
	db = initPostGresConnection()
	rdb = initRedisConnection()
}
