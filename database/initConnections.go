package database

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Rdb *redis.Client

func init() {
	Db = initPostGresConnection()
	Rdb = initRedisConnection()
}
