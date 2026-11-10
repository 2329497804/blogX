package core

import (
	"github.com/2329497804/blogX-server/global"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func InitRedis() *redis.Client {
	conf := global.Config.Redis
	redisDB := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})
	_, err := redisDB.Ping().Result()
	if err != nil {
		logrus.Fatalf("redis connect fail, %s", err)
	}
	logrus.Info("redis connect success")
	return redisDB
}
