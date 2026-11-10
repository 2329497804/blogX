package global

import (
	"github.com/2329497804/blogX-server/conf"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	Config *conf.Config
	DB     *gorm.DB
	Redis  *redis.Client
)
