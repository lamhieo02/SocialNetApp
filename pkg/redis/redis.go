package redispkg

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/lamhieo02/socialnetapp/config"
)

var (
	defaultRedisMaxActive = 0 // unlimited max active connection
	defaultRedisMaxIdle   = 10
)

func NewRedisConnection(cfg *config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB: cfg.Database,
		PoolSize: defaultRedisMaxActive,
		MinIdleConns: defaultRedisMaxIdle,
	})

	// ping to check connection
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}