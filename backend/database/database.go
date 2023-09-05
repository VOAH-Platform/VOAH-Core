package database

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var DB *gorm.DB

type RedisDB struct {
	SessionRedis       *redis.Client
	LastActivityRedis  *redis.Client
	LastRefreshRedis   *redis.Client
	PasswordResetRedis *redis.Client
	RegisterVerifyDB   *redis.Client
}

var Redis RedisDB
