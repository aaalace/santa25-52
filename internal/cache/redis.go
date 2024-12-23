package cache

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"santa25-52/internal/config"
)

func MustLoad(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		DB:   0,
	})
}
