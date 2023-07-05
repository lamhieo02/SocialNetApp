package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	redisUserEmailKey = "useremail"
)

type cache struct {
	redis *redis.Client
}

func NewCache(redis *redis.Client) *cache {
	return &cache{redis: redis}
}

func (c *cache) Set(ctx context.Context,key string, value interface{}) error {
	return c.redis.Set(ctx, redisUserEmailKey, value, 2*time.Hour).Err()
}

func (c *cache) Get(ctx context.Context, key string) (*string, error) {
	val, err := c.redis.Get(ctx, redisUserEmailKey).Result()	
	if err != nil {
		return nil, err
	}
	return &val, nil
}