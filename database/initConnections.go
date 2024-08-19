package database

import (
	"github.com/golang/glog"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Rdb *redis.Client

func init() {
	glog.Infoln("Initiating database connections...")
	DB = initPostGresConnection()
	Rdb = initRedisConnection()
}
