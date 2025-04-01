package redis

import (
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

type RedisConf struct {
	Host     string
	Password string
	DB       int
}

func New(c *RedisConf) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: c.Password, // no password set
		DB:       c.DB,       // use default DB
	})
}
