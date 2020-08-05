package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var (
	MYSQL *gorm.DB
	REDIS *redis.Client
)
