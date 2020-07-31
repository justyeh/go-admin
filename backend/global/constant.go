package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var (
	DB    *gorm.DB
	REDIS *redis.Client
)
